import json
from pathlib import Path
import networkx as nx
from networkx.readwrite import json_graph

# Load graph
with open('graphify-out/graph.json', 'r') as f:
    data = json.load(f)
G = json_graph.node_link_graph(data, edges='edges')

print(f'Graph loaded: {G.number_of_nodes()} nodes, {G.number_of_edges()} edges')

# Simple community detection using label propagation
try:
    from networkx.algorithms import community as nx_community
    communities = nx_community.label_propagation_communities(G)
    communities = {i: list(c) for i, c in enumerate(communities)}
    print(f'Found {len(communities)} communities')
except Exception as e:
    print(f'Community detection error: {e}')
    communities = {0: list(G.nodes())}

# Calculate simple cohesion (average degree within community)
cohesion = {}
for cid, members in communities.items():
    if len(members) > 1:
        subgraph = G.subgraph(members)
        avg_deg = sum(dict(subgraph.degree()).values()) / len(members)
        cohesion[cid] = avg_deg
    else:
        cohesion[cid] = 0.0

# Find god nodes (high degree, central nodes)
degrees = dict(G.degree())
sorted_nodes = sorted(degrees.items(), key=lambda x: x[1], reverse=True)
gods = []
for nid, deg in sorted_nodes[:10]:
    label = G.nodes[nid].get('label', nid)
    gods.append({'id': nid, 'label': label, 'degree': deg})

print(f'God nodes: {[(g["label"], g["degree"]) for g in gods[:5]]}')

# Find surprising connections (cross-community edges)
cross_comm = {}
node_to_comm = {}
for cid, members in communities.items():
    for n in members:
        node_to_comm[n] = cid

for u, v in G.edges():
    cu, cv = node_to_comm.get(u), node_to_comm.get(v)
    if cu is not None and cv is not None and cu != cv:
        key = (min(cu, cv), max(cu, cv))
        cross_comm[key] = cross_comm.get(key, 0) + 1

surprises = []
for (c1, c2), count in sorted(cross_comm.items(), key=lambda x: x[1], reverse=True)[:10]:
    surprises.append({
        'community_1': c1,
        'community_2': c2,
        'cross_edges': count
    })

# Generate suggested questions
questions = []
if len(communities) > 1:
    questions.append(f"How do the {len(communities)} communities relate to each other?")
if gods:
    questions.append(f"What is the role of {gods[0]['label']} in this codebase?")
questions.append("What are the main modules and their dependencies?")
questions.append("How is the client connection managed?")

# Save analysis
analysis = {
    'communities': {str(k): v for k, v in communities.items()},
    'cohesion': {str(k): v for k, v in cohesion.items()},
    'gods': gods,
    'surprises': surprises,
    'questions': questions
}
with open('graphify-out/.graphify_analysis.json', 'w') as f:
    json.dump(analysis, f, indent=2)

print('Analysis saved')
print(f'Communities: {len(communities)}')
print(f'Cohesion scores: {list(cohesion.values())[:5]}')
print(f'Surprising connections: {len(surprises)}')
print(f'Questions: {questions}')