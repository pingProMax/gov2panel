package service_relay

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
)

func init() {
	err := ReloadAllCIDRs()
	if err != nil {
		panic(fmt.Sprintf("CIDR 文件加载失败: %v", err))
	}
}

// ============ CIDR 与并发保护 ============
var (
	chinanet46 []*net.IPNet
	unicom46   []*net.IPNet
	cmcc46     []*net.IPNet
	cernet46   []*net.IPNet
	cstnet46   []*net.IPNet
	drpeng46   []*net.IPNet
	googlecn46 []*net.IPNet
	mu         sync.RWMutex
)

// ============ 文件加载和监听 ============

func ReloadAllCIDRs() error {
	var err error
	mu.Lock()
	defer mu.Unlock()

	if chinanet46, err = loadCIDRs("./resource/public/resource/china-operator-ip-ip-lists/chinanet46.txt"); err != nil {
		return fmt.Errorf("chinanet46.txt 加载失败: %v", err)
	}
	if unicom46, err = loadCIDRs("./resource/public/resource/china-operator-ip-ip-lists/unicom46.txt"); err != nil {
		return fmt.Errorf("unicom46.txt 加载失败: %v", err)
	}
	if cmcc46, err = loadCIDRs("./resource/public/resource/china-operator-ip-ip-lists/cmcc46.txt"); err != nil {
		return fmt.Errorf("cmcc46.txt 加载失败: %v", err)
	}
	if cernet46, err = loadCIDRs("./resource/public/resource/china-operator-ip-ip-lists/cernet46.txt"); err != nil {
		return fmt.Errorf("cernet46.txt 加载失败: %v", err)
	}
	if cstnet46, err = loadCIDRs("./resource/public/resource/china-operator-ip-ip-lists/cstnet46.txt"); err != nil {
		return fmt.Errorf("cstnet46.txt 加载失败: %v", err)
	}
	if drpeng46, err = loadCIDRs("./resource/public/resource/china-operator-ip-ip-lists/drpeng46.txt"); err != nil {
		return fmt.Errorf("drpeng46.txt 加载失败: %v", err)
	}
	if googlecn46, err = loadCIDRs("./resource/public/resource/china-operator-ip-ip-lists/googlecn46.txt"); err != nil {
		return fmt.Errorf("googlecn46.txt 加载失败: %v", err)
	}

	log.Println("CIDR 文件加载完成")
	fmt.Println("chinanet46: ", len(chinanet46))
	fmt.Println("unicom46: ", len(unicom46))
	fmt.Println("cmcc46: ", len(cmcc46))
	fmt.Println("cernet46: ", len(cernet46))
	fmt.Println("cstnet46: ", len(cstnet46))
	fmt.Println("drpeng46: ", len(drpeng46))
	fmt.Println("googlecn46: ", len(googlecn46))
	fmt.Println()
	return nil
}

func WatchCIDRFiles() {
	files := []string{
		"./resource/public/resource/china-operator-ip-ip-lists/chinanet46.txt",
		"./resource/public/resource/china-operator-ip-ip-lists/unicom46.txt",
		"./resource/public/resource/china-operator-ip-ip-lists/cmcc46.txt",
		"./resource/public/resource/china-operator-ip-ip-lists/cernet46.txt",
		"./resource/public/resource/china-operator-ip-ip-lists/cstnet46.txt",
		"./resource/public/resource/china-operator-ip-ip-lists/drpeng46.txt",
		"./resource/public/resource/china-operator-ip-ip-lists/googlecn46.txt",
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("创建文件监听器失败:", err)
	}
	defer watcher.Close()

	for _, f := range files {
		if err := watcher.Add(f); err != nil {
			log.Printf("监听文件失败: %s, %v\n", f, err)
		}
	}

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Printf("文件修改，重新加载: %s\n", event.Name)
				if err := ReloadAllCIDRs(); err != nil {
					log.Println("刷新 CIDR 文件失败:", err)
				}
			}
		case err := <-watcher.Errors:
			log.Println("文件监听错误:", err)
		}
	}
}

// ============ ASN 判断 ============
func GetASN(ipStr string) string {
	mu.RLock()
	defer mu.RUnlock()

	ip := net.ParseIP(ipStr)
	if ip == nil {
		return ""
	}

	switch {
	case ipInCIDRs(ip, chinanet46):
		return "AS4134"
	case ipInCIDRs(ip, unicom46):
		return "AS4837"
	case ipInCIDRs(ip, cmcc46):
		return "AS9808"
	// case ipInCIDRs(ip, cernet46):
	// 	return "AS4538"
	// case ipInCIDRs(ip, cstnet46):
	// 	return "AS4539"
	// case ipInCIDRs(ip, drpeng46):
	// 	return "AS4540"
	// case ipInCIDRs(ip, googlecn46):
	// 	return "AS4541"
	default:
		return ""
	}
}

// ============ 辅助函数 ============

func loadCIDRs(filename string) ([]*net.IPNet, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cidrs []*net.IPNet
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		_, ipnet, err := net.ParseCIDR(line)
		if err != nil {
			fmt.Printf("忽略无效CIDR: %s\n", line)
			continue
		}
		cidrs = append(cidrs, ipnet)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return cidrs, nil
}

func ipInCIDRs(ip net.IP, cidrs []*net.IPNet) bool {
	for _, c := range cidrs {
		if c.Contains(ip) {
			return true
		}
	}
	return false
}
