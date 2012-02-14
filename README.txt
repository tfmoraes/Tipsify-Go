This is a implementation of algorithm Tipsify from "Fast Triangle
Reordering for Vertex Locality and Reduced Overdraw". This algorithm is to sort
a given mesh, which implies in an reduction of time to render that mesh. The
idea of this algorithm is to explore the locality in mesh and the cache from
Video Card.

To Compile is only necessary to have the Go's compiler installed. Then:

$ make

To execute:

$ ./tipsify_sorter input_mesh.ply output_mesh.ply

