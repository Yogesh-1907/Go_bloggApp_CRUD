package main

import (
	"fmt"
	"gorm_blogging_application/repository"
	"gorm_blogging_application/service"
	"gorm_blogging_application/utility"
)

var (
	postService = service.NewPostServiceImpl(repository.PostRepoImpl{})
)

func main() {
	utility.DbConnect()
	//getAllPosts()
	// getPostsPaginated()
	// getPostsPaginatedByTitle()
	getPostsPaginatedByAuthorID()
}

// to retrieve list of all posts from database
func getAllPosts() {
	posts, err := postService.GetAllPosts()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if len(posts) == 0 {
			fmt.Println("posts not found")
		} else {
			fmt.Println("All Posts:")
			for _, post := range posts {
				fmt.Println("ID:", post.ID)
				fmt.Println("Title:", post.Title)
				fmt.Println("Content:", post.Content)
				fmt.Println("Author:", post.Author)
				fmt.Println()
			}
		}
	}
}

// to retrieve list of posts from database that are paginated
func getPostsPaginated() {
	pageNo := 1
	pageSize := 5
	posts, err := postService.GetPostsPaginated(pageNo, pageSize)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if len(posts) == 0 {
			fmt.Println("posts not found")
		} else {
			fmt.Printf("Showing %d posts from page %d \n", pageSize, pageNo)
			for _, post := range posts {
				fmt.Println("ID:", post.ID)
				fmt.Println("Title:", post.Title)
				fmt.Println("Content:", post.Content)
				fmt.Println("Author:", post.Author)
				fmt.Println()
			}
		}
	}
}

// to retrieve list of posts from database that are paginated by title
func getPostsPaginatedByTitle() {
	pageNo := 1
	pageSize := 5
	posts, err := postService.GetPostsPaginatedByTitle(pageNo, pageSize)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if len(posts) == 0 {
			fmt.Println("posts not found")
		} else {
			fmt.Printf("Showing %d posts from page %d \n", pageSize, pageNo)
			for _, post := range posts {
				fmt.Println("ID:", post.ID)
				fmt.Println("Title:", post.Title)
				fmt.Println("Content:", post.Content)
				fmt.Println("Author:", post.Author)
				fmt.Println()
			}
		}
	}
}

// to retrieve list of posts from database that are paginated by authorid in descending order
func getPostsPaginatedByAuthorID() {
	pageNo := 1
	pageSize := 5
	posts, err := postService.GetPostsPaginatedByAuthorID(pageNo, pageSize)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if len(posts) == 0 {
			fmt.Println("posts not found")
		} else {
			fmt.Printf("Showing %d posts from page %d \n", pageSize, pageNo)
			for _, post := range posts {
				fmt.Println("ID:", post.ID)
				fmt.Println("Title:", post.Title)
				fmt.Println("Content:", post.Content)
				fmt.Println("Author:", post.Author)
				fmt.Println()
			}
		}
	}
}
