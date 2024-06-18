package maps

import (
	"crypto/md5"
	"fmt"
	"io"
)

func getCounts(userIDs []string) map[string]int {
	counts := make(map[string]int)
	for _, userID := range userIDs {
		// 값이 없으니 0 반환 
		count := counts[userID]
		count++
		counts[userID] = count
	}

	return counts
}


func countsMapTest(userIDs []string, ids []string) {
	fmt.Printf("Generating counts for %v user IDs...\n", len(userIDs))

	counts := getCounts(userIDs)
	fmt.Println("Counts from select IDs:")
	for _, k := range ids {
		v := counts[k]
		fmt.Printf(" - %s: %d\n", k, v)
	}
	fmt.Println("=====================================")
}

func countsMap() {
	userIDs := []string{}
	for i := 0; i < 10000; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprint(i))
		key := fmt.Sprintf("%x", h.Sum(nil))
		userIDs = append(userIDs, key[:2])
	}

	countsMapTest(userIDs, []string{"00", "ff", "dd"})
	countsMapTest(userIDs, []string{"aa", "12", "32"})
	countsMapTest(userIDs, []string{"bb", "33"})
}