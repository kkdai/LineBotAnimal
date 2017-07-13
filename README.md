ASKME Animal Bot: A Animal Classification Chatbot using Tensorflow inception model
==============

[![Join the chat at https://gitter.im/kkdai/LineBotAnimal](https://badges.gitter.im/kkdai/LineBotAnimal.svg)](https://gitter.im/kkdai/LineBotAnimal?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

 [![GoDoc](https://godoc.org/github.com/kkdai/LineBotAnimal.svg?status.svg)](https://godoc.org/github.com/kkdai/LineBotAnimal)  [![Build Status](https://travis-ci.org/kkdai/LineBotAnimal.svg?branch=master)](https://travis-ci.org/kkdai/LineBotAnimal.svg)

[![goreportcard.com](https://goreportcard.com/badge/github.com/kkdai/LineBotAnimal)](https://goreportcard.com/report/github.com/kkdai/LineBotAnimal)


![](images/icon.PNG)


## Features:

- Upload an animal photo, this chatbot will tell you what animal it is.
- All the training model using [Tensorflow Inception](https://github.com/tensorflow/models/tree/master/inception) result.

## Just Deploy the same on Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)


How to use it
---------------

### Add friend:

![](images/qrcode.png)

[![加入好友](https://scdn.line-apps.com/n/line_add_friends/btn/zh-Hant.png)](https://line.me/R/ti/p/%40ujo0893j)

- Upload photo 
- See result



![](images/how_use.PNG)


How to build your own Tensorflow Chatbot in Golang
---------------

This chatbot will link to an API server which have tensforflow prebuild model file.

- Refer [https://github.com/kkdai/tf-go-inception](https://github.com/kkdai/tf-go-inception) for how to build tensorflow prebuild model API server. 
- Add in Heroko `Config Variables` with
	- `ApiURL=http://1.2.3.4:3000/api/v1/tf-image/`

License
---------------

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

