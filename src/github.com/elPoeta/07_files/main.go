package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	isWriteFile := writeContent("myFile.txt", getStaticContent())
	if isWriteFile {
		fmt.Println("File was created!")
	}
	fmt.Println(readContent("myFile.txt"))
}

func writeContent(fileName string, content string) bool {
	err := ioutil.WriteFile(fileName, []byte(content), 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return true
}

func readContent(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(content)
}

func getStaticContent() string {
	return "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec accumsan rhoncus lacus, non finibus urna pharetra vel. Maecenas tempus purus id vestibulum blandit. Cras consequat, risus vitae tristique accumsan, dolor lacus scelerisque nisl, at dignissim ipsum neque vel ante. Nulla nec dignissim neque. Sed volutpat massa quis mi convallis vulputate. Vivamus in lectus facilisis felis feugiat volutpat sed vel turpis. Pellentesque magna neque, convallis at sapien in, pulvinar vehicula neque. Cras erat sapien, sollicitudin a purus sit amet, pharetra semper urna. In hac habitasse platea dictumst. Aliquam facilisis aliquet nulla eget porta.\nOrci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Aliquam nec nulla quis magna blandit faucibus. In hac habitasse platea dictumst. Morbi dignissim mattis neque, semper mollis turpis vestibulum eget. Nullam quis dolor risus. Vivamus sodales lorem dolor. Suspendisse quis purus vel justo ullamcorper rutrum non vel ligula. Cras nec tincidunt tellus. Suspendisse diam erat, placerat bibendum vulputate eu, eleifend in ex. Nam et suscipit enim. Nullam magna risus, luctus sit amet nibh et, euismod fringilla felis. Aliquam lobortis tortor vel scelerisque mollis. Sed porttitor commodo ex, dictum dictum nisl aliquet nec. Nunc ultrices sem non metus hendrerit, porta tempor ligula hendrerit. Aenean viverra in erat eu volutpat. Suspendisse auctor orci ut augue bibendum facilisis.\nNullam vehicula urna sagittis, rhoncus dolor et, luctus lacus. Sed sed felis sem. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque consequat mollis nulla, vel laoreet ipsum consectetur non. Nulla in elementum justo, at scelerisque magna. Suspendisse id lectus tincidunt, mollis tortor in, hendrerit justo. Quisque a diam pellentesque, aliquet nunc at, interdum sapien. Pellentesque eu elit et dolor ultricies aliquam. Suspendisse semper orci finibus dictum mollis.\nMorbi augue est, vulputate mollis vulputate tempus, efficitur sit amet lorem. Fusce molestie, velit in euismod ornare, dui ex pulvinar mi, quis laoreet lectus nisl nec lorem. Duis erat lacus, malesuada at hendrerit vitae, maximus in sem. Phasellus finibus orci risus, et porta mauris accumsan nec. In sit amet nisl diam. Maecenas eu faucibus eros. Vivamus tempor sapien sit amet mattis pretium. Aliquam eu massa non sem vulputate sagittis. Duis dui dolor, pretium tempus porta eget, rhoncus eget mi. Aliquam dictum mi id est finibus congue. Nunc sit amet convallis lacus. Nullam mollis risus at elementum gravida. Phasellus non rhoncus metus. Duis arcu enim, porta quis diam eu, tristique viverra ligula.\nAliquam nec nisl accumsan, blandit felis vel, efficitur nisi. In hac habitasse platea dictumst. Ut sagittis tincidunt nibh. Suspendisse dictum placerat neque, vitae malesuada nisl sagittis sed. Quisque venenatis dictum magna ac dignissim. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae"
}
