// https://leetcode.com/problems/design-twitter/

package main

import (
	"container/heap"
	"fmt"
	"time"
)

type tweet struct {
	ID       int
	postTime int64
	next     *tweet
}
type ll struct {
	head *tweet
}

func (l *ll) add(n *tweet) {
	if l.head == nil {
		l.head = n
	} else {
		n.next = l.head
		l.head = n
	}
}

type MaxHeap []tweet

func (mh MaxHeap) Less(i, j int) bool {
	return mh[i].postTime > mh[j].postTime
}
func (mh MaxHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}
func (mh MaxHeap) Len() int {
	return len(mh)
}
func (mh *MaxHeap) Pop() any {
	n := mh.Len() - 1
	v := (*mh)[n]
	(*mh) = (*mh)[:n]
	return v
}
func (mh *MaxHeap) Push(v any) {
	(*mh) = append((*mh), v.(tweet))
}

type User struct {
	following map[int]bool
	tweets    *ll
}

func mergeKLists(allTweets []*tweet) []int {
	if len(allTweets) == 0 {
		return nil
	}
	mh := MaxHeap{}
	numTweets := 10
	ans := []int{}

	for _, t := range allTweets {
		heap.Push(&mh, *t)
	}

	for mh.Len() != 0 {
		if len(ans) == numTweets {
			break
		}
		t := heap.Pop(&mh).(tweet)
		if t.next != nil {
			heap.Push(&mh, *t.next)
		}
		ans = append(ans, t.ID)
	}
	return ans
}

type Twitter struct {
	users map[int]User
}

func Constructor() Twitter {
	return Twitter{
		users: map[int]User{},
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, found := this.users[userId]; !found {
		this.users[userId] = User{
			following: map[int]bool{},
			tweets:    &ll{},
		}
	}
	// Post for user
	this.users[userId].tweets.add(&tweet{
		ID:       tweetId,
		postTime: time.Now().UnixNano(),
	})
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	ans := []int{}
	if _, found := this.users[userId]; !found {
		return ans
	}
	allTweets := []*tweet{}
	if this.users[userId].tweets != nil {
		if this.users[userId].tweets.head != nil {
			allTweets = append(allTweets, this.users[userId].tweets.head)
		}
	}
	for follows := range this.users[userId].following {
		if this.users[follows].tweets != nil {
			if this.users[follows].tweets.head != nil {
				allTweets = append(allTweets, this.users[follows].tweets.head)
			}
		}
	}
	return mergeKLists(allTweets)
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, found := this.users[followerId]; !found {
		this.users[followerId] = User{
			following: map[int]bool{},
			tweets:    &ll{},
		}
	}
	this.users[followerId].following[followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.users[followerId].following, followeeId)
}

func main() {
	// obj := Constructor()
	// obj.PostTweet(1, 5)
	// obj.PostTweet(1, 50)
	// obj.PostTweet(1, 51)
	// fmt.Println(obj.users)
	// fmt.Println(obj.GetNewsFeed(1))
	// obj.Follow(1, 2)
	// fmt.Println(obj.users)
	// obj.PostTweet(2, 6)
	// fmt.Println(obj.users)
	// fmt.Println(obj.GetNewsFeed(1))
	// obj.Unfollow(1, 2)
	// fmt.Println(obj.users)
	// fmt.Println(obj.GetNewsFeed(1))

	// obj.Follow(1, 2)
	// obj.PostTweet(2, 60)
	// obj.PostTweet(2, 61)
	// obj.PostTweet(2, 63)
	// obj.PostTweet(2, 64)
	// obj.PostTweet(2, 65)
	// obj.PostTweet(2, 66)
	// obj.PostTweet(2, 67)
	// obj.PostTweet(2, 68)
	// fmt.Println(obj.GetNewsFeed(1))

	obj1 := Constructor()
	obj1.PostTweet(1, 1)
	fmt.Println(obj1.GetNewsFeed(1))
	obj1.Follow(2, 1)
	fmt.Printf("%+v\n", obj1.users)
	fmt.Println(obj1.GetNewsFeed(2))

	obj2 := Constructor()
	obj2.PostTweet(1, 5)
	fmt.Println(obj2.GetNewsFeed(1))
	obj2.Follow(2, 1)
	obj2.Follow(1, 2)
	fmt.Printf("%+v\n", obj2.users)
	fmt.Println(obj2.GetNewsFeed(2))
	obj2.PostTweet(2, 6)
	fmt.Printf("%+v\n", obj2.users)

}
