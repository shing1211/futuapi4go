import json
import sys
import os
from pathlib import Path

# Set environment BEFORE importing anything that might use multiprocessing
os.environ['PYTHON_MULTIPROCESSING'] = '0'

# Try to use graphify's single-threaded mode if available
try:
    import graphify
    # Check if there's a single-threaded option
    if hasattr(graphify.extract, 'extract_sequential'):
        from graphify.extract import extract_sequential as extract
        print('Using sequential extraction')
    else:
        from graphify.extract import extract
except Exception as e:
    print(f'Graphify import issue: {e}')

# Since graphify has Windows multiprocessing issues, let's do manual Go AST extraction
# This will parse the Go files and extract relationships manually

print('Performing manual Go code extraction...')

# Read the detect result to get the code files
detect_file = Path('graphify-out/.graphify_detect.json')
with open(detect_file, 'r') as f:
    detect = json.load(f)

code_files = detect.get('files', {}).get('code', [])
print(f'Found {len(code_files)} code files')

# Simple extraction - read each file and extract import/function relationships
import re

nodes = []
edges = []
node_id = 0

def make_id(name, file_type):
    return f"{name}_{node_id}"

# Parse Go files for imports and functions
for filepath in code_files:
    try:
        with open(filepath, 'r', encoding='utf-8', errors='ignore') as f:
            content = f.read()
        
        # Skip proto files - they're generated
        if filepath.endswith('.pb.go'):
            continue
            
        # Get module name from file path
        rel_path = filepath.replace('D:\\github\\futuapi4go\\', '').replace('\\', '/')
        
        # Extract package
        pkg_match = re.search(r'^package\s+(\w+)', content, re.MULTILINE)
        pkg = pkg_match.group(1) if pkg_match else 'unknown'
        
        # Create module node
        safe_path = rel_path.replace('.go','').replace('/','_')
        module_id = f"module_{safe_path}"
        nodes.append({
            'id': module_id,
            'label': pkg,
            'file_type': 'code',
            'source_file': rel_path,
            'source_location': None
        })
        
        # Extract imports
        import_block = re.search(r'import\s*\(([^)]*)\)|import\s+"([^"]+)"', content, re.MULTILINE)
        if import_block:
            imports = re.findall(r'"([^"]+)"', import_block.group(0))
            for imp in imports:
                if not imp.startswith('github.com/shing1211/futuapi4go'):
                    continue
                imp_id = f"import_{imp.replace('github.com/shing1211.futuapi4go.','').replace('/','_')}"
                nodes.append({
                    'id': imp_id,
                    'label': imp.split('/')[-1],
                    'file_type': 'code',
                    'source_file': rel_path,
                    'source_location': None
                })
                edges.append({
                    'source': module_id,
                    'target': imp_id,
                    'relation': 'imports',
                    'confidence': 'EXTRACTED',
                    'confidence_score': 1.0,
                    'source_file': rel_path,
                    'weight': 1.0
                })
        
        # Extract function definitions
        func_matches = re.findall(r'func\s+(?:\(\w+\s+\*?\w+\)\s+)?(\w+)\s*\(', content)
        for func in func_matches:
            if func[0].isupper() and len(func) > 1:  # Public functions only
                safe_path = rel_path.replace('.go','').replace('/','_')
                func_id = f"func_{func}_{safe_path}"
                nodes.append({
                    'id': func_id,
                    'label': func,
                    'file_type': 'code',
                    'source_file': rel_path,
                    'source_location': None
                })
                edges.append({
                    'source': module_id,
                    'target': func_id,
                    'relation': 'defines',
                    'confidence': 'EXTRACTED',
                    'confidence_score': 1.0,
                    'source_file': rel_path,
                    'weight': 1.0
                })
        
        # Extract type definitions
        type_matches = re.findall(r'type\s+(\w+)\s+struct', content)
        for t in type_matches:
            safe_path = rel_path.replace('.go','').replace('/','_')
            type_id = f"type_{t}_{safe_path}"
            nodes.append({
                'id': type_id,
                'label': t,
                'file_type': 'code',
                'source_file': rel_path,
                'source_location': None
            })
            edges.append({
                'source': module_id,
                'target': type_id,
                'relation': 'defines',
                'confidence': 'EXTRACTED',
                'confidence_score': 1.0,
                'source_file': rel_path,
                'weight': 1.0
            })
                
    except Exception as e:
        print(f'Error processing {filepath}: {e}')
        continue

result = {'nodes': nodes, 'edges': edges, 'input_tokens': 0, 'output_tokens': 0}

Path('graphify-out/.graphify_ast.json').write_text(json.dumps(result, indent=2))
print(f'Manual extraction: {len(result["nodes"])} nodes, {len(result["edges"])} edges')

# Now build the graph using networkx directly
import networkx as nx
from networkx.readwrite import json_graph

G = nx.DiGraph()
for node in nodes:
    G.add_node(node['id'], **node)
for edge in edges:
    G.add_edge(edge['source'], edge['target'], **edge)

print(f'Graph: {G.number_of_nodes()} nodes, {G.number_of_edges()} edges')

# Save as graph.json
graph_data = json_graph.node_link_data(G)
with open('graphify-out/graph.json', 'w') as f:
    json.dump(graph_data, f, indent=2)

print('graph.json written')