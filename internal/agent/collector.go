package agent

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/amir/sys-monitor/proto"
)

func CollectMetrics(serverID string) (*proto.MetricRequest, error) {

	cpuUsage, err := getCPUUsage()
	if err != nil {
		return nil, fmt.Errorf("Ошибка сбора cpu %v", err)
	}

	freeDisk, err := getFreeDiskSpace("/")
	if err != nil {
		return nil, fmt.Errorf("Ошибка при сборе свободного пространства %v", err)
	}

	return &proto.MetricRequest{
		ServerId:      serverID,
		CpuUsage:      cpuUsage,
		FreeDiskSpace: uint64(freeDisk),
		Timestamp:     time.Now().Unix(),
	}, nil

}

func getCPUUsage() (float32, error) {

	data, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return 0, err
	}

	infoArray := strings.Fields(string(data))
	if len(infoArray) == 0 {
		return 0, errors.New("Пустой файл /proc/loadavg")
	}

	cpu, err := strconv.ParseFloat(infoArray[0], 32)
	if err != nil {
		return 0, err
	}

	return float32(cpu), nil

}

func getFreeDiskSpace(path string) (int64, error) {

	var stat syscall.Statfs_t
	err := syscall.Statfs(path, &stat)
	if err != nil {
		return 0, err
	}

	freeSpaceByte := stat.Bavail * uint64(stat.Bsize)
	return int64(freeSpaceByte), nil

}
