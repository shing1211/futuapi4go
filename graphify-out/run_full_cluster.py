import json
from pathlib import Path
import networkx as nx
from networkx.readwrite import json_graph

# Load graph
with open('graphify-out/graph.json', 'r') as f:
    data = json.load(f)
G = json_graph.node_link_graph(data, edges='edges')

print(f'Graph: {G.number_of_nodes()} nodes, {G.number_of_edges()} edges')

# Convert to undirected for community detection
G_undirected = G.to_undirected()

# Community detection
from networkx.algorithms import community as nx_community
communities = nx_community.label_propagation_communities(G_undirected)
communities = {i: list(c) for i, c in enumerate(communities)}
print(f'Found {len(communities)} communities')

# Calculate cohesion
cohesion = {}
for cid, members in communities.items():
    if len(members) > 1:
        subgraph = G_undirected.subgraph(members)
        avg_deg = sum(dict(subgraph.degree()).values()) / len(members)
        cohesion[cid] = round(avg_deg, 2)
    else:
        cohesion[cid] = 0.0

# Find god nodes
degrees = dict(G.degree())
sorted_nodes = sorted(degrees.items(), key=lambda x: x[1], reverse=True)
gods = []
for nid, deg in sorted_nodes[:10]:
    label = G.nodes[nid].get('label', nid)
    gods.append({'id': nid, 'label': label, 'degree': deg})

# Cross-community edges
node_to_comm = {}
for cid, members in communities.items():
    for n in members:
        node_to_comm[n] = cid

cross_comm = {}
for u, v in G.edges():
    cu, cv = node_to_comm.get(u), node_to_comm.get(v)
    if cu is not None and cv is not None and cu != cv:
        key = (min(cu, cv), max(cu, cv))
        cross_comm[key] = cross_comm.get(key, 0) + 1

surprises = []
for (c1, c2), count in sorted(cross_comm.items(), key=lambda x: x[1], reverse=True)[:10]:
    l1 = communities.get(c1, ['unknown'])[:3]
    l2 = communities.get(c2, ['unknown'])[:3]
    surprises.append({
        'community_1': c1,
        'community_2': c2,
        'cross_edges': count,
        'sample_nodes_1': [G.nodes[n].get('label',n) for n in l1[:2]],
        'sample_nodes_2': [G.nodes[n].get('label',n) for n in l2[:2]]
    })

# Auto-label communities based on most common package prefixes
labels = {}
for cid, members in communities.items():
    node_labels = [G.nodes[n].get('label', '') for n in members]
    # Find common package prefixes
    prefixes = {}
    for label in node_labels:
        if '_' in label:
            prefix = label.split('_')[0]
            prefixes[prefix] = prefixes.get(prefix, 0) + 1
    if prefixes:
        top_prefix = max(prefixes.keys(), key=lambda k: prefixes[k])
        labels[cid] = top_prefix.capitalize() + " Module"
    else:
        labels[cid] = f"Community {cid}"

print(f'Labels: {labels}')

# Suggested questions
questions = [
    f"What is the role of {gods[0]['label']} in this API client?",
    "What are the main modules and their dependencies?",
    "How does the client connection management work?",
    "What trading (trd) APIs are available?",
    "What market data (qot) features does this library provide?"
]

# Save analysis
analysis = {
    'communities': {str(k): v for k, v in communities.items()},
    'cohesion': {str(k): v for k, v in cohesion.items()},
    'gods': gods,
    'surprises': surprises,
    'questions': questions,
    'labels': {str(k): v for k, v in labels.items()}
}
with open('graphify-out/.graphify_analysis.json', 'w') as f:
    json.dump(analysis, f, indent=2)

# Save labels separately
with open('graphify-out/.graphify_labels.json', 'w') as f:
    json.dump({str(k): v for k, v in labels.items()}, f)

print(f'Communities: {len(communities)}')
print(f'Cohesion: {cohesion}')
print(f'Questions: {questions}')
print('Done!')