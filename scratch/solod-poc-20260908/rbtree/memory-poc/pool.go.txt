package kernels

// Go owns typed, fixed-size blocks. Appending block pointers never moves nodes.
// The pool and its nodes are intentionally confined to one benchmark worker.
const blockNodes = 32768 // 1 MiB: each RBTree node is 32 bytes on this platform.
type nodeBlock [blockNodes]Constructor_Test_RBTree_T
var blocks []*nodeBlock
var currentBlock *nodeBlock
var currentIndex = blockNodes
var nextBlock int
var maxBlocks int
var pooledNodes int64
var fallbackNodes int64

func ConfigurePool(capMiB int) { maxBlocks = capMiB }

// Called only after every reference to the previous tree has died.
// Clearing is included in benchmarks: stale pointers must not retain heap data.
func ResetPool() {
    remaining := pooledNodes
    for _, block := range blocks {
        if remaining == 0 { break }
        used := int64(blockNodes)
        if remaining < used { used = remaining }
        clear(block[:int(used)])
        remaining -= used
    }
    currentBlock = nil
    currentIndex = blockNodes
    nextBlock = 0
    pooledNodes = 0
    fallbackNodes = 0
}

func AllocNode(rc uint32, color uint32, left *Constructor_Test_RBTree_T, key int64, right *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
    if currentIndex >= blockNodes {
        if nextBlock >= maxBlocks {
            fallbackNodes++
            return &Constructor_Test_RBTree_T{rc, color, left, key, right}
        }
        if nextBlock == len(blocks) { blocks = append(blocks, new(nodeBlock)) }
        currentBlock = blocks[nextBlock]
        nextBlock++
        currentIndex = 0
    }
    p := &currentBlock[currentIndex]
    currentIndex++
    pooledNodes++
    *p = Constructor_Test_RBTree_T{rc, color, left, key, right}
    return p
}

func AllocationCounts() (int64, int64, int64) {
    return pooledNodes+fallbackNodes, int64(len(blocks))*blockNodes*32, fallbackNodes
}
