# Graph

# representation
https://www.khanacademy.org/computing/computer-science/algorithms/graph-representation/a/representing-graphs
1) Edge list:
   array of edges. each element in the array has two values. source node -> dest node. the pair [0,1] indicates that there is and edge from 0 to 1
2) Adjencency list. 
  An array or map is used to represent adjencecy list. An adjency is define as source node and all nodes that are connected with that node. 
  - An array represenation:
    each index in the array is repesents the node number. and the list indicates the neightbors it is connecting. 
    [1,2,3]  // 0 node connects to 1,2,3
    [3,4] // 1 node connects to 3,4
    [] // 2 node connects to no one
    [] //3 node connects to no one.
  - map reperesentation: 
    key is the address of a node and value is the list of nodes it is connecting.
    ```go
    map[*node][]*node
    ```
  - structure representation:
    ```go
    graph = []*node
    type node struct{
        val int
        neighbors []*node
    }
    ```
Pros:
  - Does not take extra space. if there is V node and E edges, it will take O(E) space
  - finding neighbors(children) is faster as you dont need to visit all V nodes to find it out.
Cons:
  - to find out if X node is connected to Y, it could take O(n) time. since if we start with root that is not X and Y then in connected graph it can take O(V) and disconnected graph it will also take O(V) time.

3) Matrix representation.
   representated by VXV size matrix. where V is no of nodes in the graph.
   if value of matrix[i][j]=1 (or some weight) that means there is an eddge between i and j
   if value of matrix[i][j]=0 (or nil) that means there is no edge between i and j
Pros:
  If we want to find if there is any edge between i and j it takes O(1) time to find out.

Cons:
  - it takes O(v^2) space. even there is very few edges between nodes. (most of the cells have value of 0)
  - if you want to find out all connecting node to i then you have to scan all the elements in that row. so even if you only have one edge you have to scan V columns


# Properties
 - traversal DFS, BFS, topological traversal.
## DFS
  Visit root, mark root as visited and then visit its children recursively.  ex. [1]->[2,3,4], [2]->[3,4] then visit 1 than visit 2 and all child of 2 than all child of 3 and than all child of 4
  - Running time O(V+E), space O(V) since extra visited list of size V is required.
  - Handle disconnected graph
    - If there are multiple graphs that are disconnected then first pick up any graph and do DFS. then go through all vertices and find if it is in visited node or not.
    if it is not in visited node, do DFS on that vertices.

  - Applications
    -  For a *weighted* graph, DFS traversal of the graph produces the minimum spanning tree and all pair shortest path tree.
       - why only weighted graph: if each children might have different weights than dpending on weight requirement, we might want to visit a children with least or most weight. 
       - This if it is unweighted graph then DFS is useless as we can go to any neighbors. BFS makes better sense in that case

    - Detecting a cycle in Graph. easy to back track to see if node is visited or not.
    -  Path finding: 
       DFS algorithm to find a path between two given vertices u and z.
        i) Call DFS(G, u) with u as the start vertex.
        ii) Use a stack S to keep track of the path between the start vertex and the current vertex.
        iii) As soon as destination vertex z is encountered, return the path as the
        contents of the stack
    - Topological sorting
      - Topological Sorting is mainly used for scheduling jobs from the given dependencies among jobs. 
      - scheduling, ordering of formula cell evaluation when recomputing formula values in spreadsheets, logic synthesis, determining the order of compilation tasks to perform in makefiles, data serialization, and resolving symbol dependencies in linkers [2].
    - bipartile graph.
      - if we dvidie graph in to two, each edge in one graph has a one node in each graphs.
      - We can augment either BFS or DFS when we first discover a new vertex, color it opposited its parents, and for each other edge, check it doesn’t link two vertices of the same color. 
      
## Topological sorting
  Topological sorting for Directed Acyclic Graph (DAG) is a linear ordering of vertices such that for every directed edge uv, vertex u comes before v in the ordering. Topological Sorting for a graph is not possible if the graph is not a DAG.
  - Modified DFS: In dfs we visit node, mark it as visited then visit its neighbors. In topoloical sort, we mark node as visited only after its all neighbors and their neighbors are visited.
  - ex. build dependency, course taken, alien dictionary problems. meeting rooms


## BFS
 Visit node, mark it as visited then visit all neighbors first than neighbors of each neighbors.
  - Running time O(V+E)
  - Applications
   - Shortest Path and Minimum Spanning Tree for unweighted graph:
     -  for *unweighted* graph all edge has same weight. so minimum path between two nodes is number of minimum edges.
     - Peer to peer network.
     - crawlers and search engine. 
       - Crawlers build index using Breadth First. The idea is to start from source page and follow all links from source and keep doing same. Depth First Traversal can also be used for crawlers, but the advantage with Breadth First Traversal is, depth or levels of the built tree can be limited.
     -  Social Networking Websites : to reperesen relation between two persons.
     - GPS - find coordinates.
     - Cycle detection in undirected graph.
     - Finding all nodes within one connected component

# Bipartile Graph.
  - https://www.geeksforgeeks.org/bipartite-graph/
  A Bipartite Graph is a graph whose vertices can be divided into two independent sets, U and V such that every edge (u, v) either connects a vertex from U to V or a vertex from V to U. In other words, for every edge (u, v), either u belongs to U and v to V, or u belongs to V and v to U. We can also say that there is no edge that connects vertices of same set.
  - 1. Assign RED color to the source vertex (putting into set U).
    2. Color all the neighbors with BLUE color (putting into set V).
    3. Color all neighbor’s neighbor with RED color (putting into set U).
    4. This way, assign color to all vertices such that it satisfies all the constraints of m way coloring problem where m = 2.
    5. While assigning colors, if we find a neighbor which is colored with same color as current vertex, then the graph cannot be colored with 2 vertices (or graph is not Bipartite)
 # Intresting problems
 - Build order depedency
 - minimum distance between two points
 - find a node which has minimum distance to all target node given obstacles. ex. Shortest Distance from All Buildings (https://leetcode.com/problems/shortest-distance-from-all-buildings/)
