### 测试命令

#### 1、Get
`curl localhost:8080/users`

#### 1、Post
`curl -H "Content-Type: application/json" -X POST -d '{"id": 666, "name":"fudute" }' localhost:8080/users`