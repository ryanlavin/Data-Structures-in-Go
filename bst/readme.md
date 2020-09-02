# Binary Search Tree
### Written using GoLang

###### Exported functions include:

**Insert** takes any struct implementing the *bstNode* interface as a parameter, initializes a pointer to a *node* struct with the *bstNode* as a field, and then iteratively sorts through the binary search tree to find the correct location to place the newly-created node.

**Search** takes any struct implementing the *bstNode* interface as a parameter and proceeds to traverse through the binary search tree until either a node with a value equal to the parameter's value (in which case a pointer to that node is returned) or the end of the tree is reached (in which case nil is returned).

**Remove** takes any struct implementing the *bstNode* interface as a parameter and first utilizes **Search** to ensure that a node with a value equal to that of the parameter exists. If so, the node is removed: the process in which the node is removed depends on whether the node to remove is a leaf node, has exactly one child, or has two children. 
