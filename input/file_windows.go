package input

import "os"

func fileIds(info *os.FileInfo) (uint64, uint64) {
	// No dev and inode numbers on windows, right?
	return 0, 0
}
