## 缓存是什么？

​	在软件系统中使用缓存,可以降低系统响应时间,提高用户体验,降低某些系统模块的压力.

​	提到缓存，可能第一时间想到的就是Redis、Memcache等，这些都属于是分布式缓存，而在某些场景下我们可能并不需要分布式缓存，毕竟需要多引入维护一个中间件，那么在数据量小，且访问频繁，或者说一些不会变的静态配置数据我们都可以考虑放置到本地缓存中。

​	谈起缓存，那不得不提**缓存淘汰策略**、**缓存过期策略**等下面是几个经典的go cache库：

## 常用缓存库

### goburrow/cache

地址：[goburrow/cache: Mango Cache 🥭 - Partial implementation of Guava Cache in Go (golang). (github.com)](https://github.com/goburrow/cache)

Go 中 Guava Cache 的部分实现。

[guava cache详细介绍_赶路人儿的博客-CSDN博客_guavacache](https://blog.csdn.net/liuxiao723846/article/details/108392072)

[groupcache入门_地鼠工程师的博客-CSDN博客_groupcache](https://blog.csdn.net/m0_37731056/article/details/104782685)

### groupcache（11.6k star）

groupcache是一款开源的缓存组件。与memcache与redis不同的是，groupcache不需要单独的部署，可以作为你程序的一个库来使用.。这样方便我们开发的程序部署。

[轻量级Golang分布式缓存groupcache笔记 - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/53724255/)

下面是几个使用示例：

[groupcache 使用示例 - 简书 (jianshu.com)](https://www.jianshu.com/p/f0ef2664428e?u_atoken=5b653377-2107-4c65-bb32-2ef989533297&u_asession=014shsMip5k0Co1FI9xyoLcfLQEHO2asP1VDHqsqE1vosewxadYqh1K0oe42K0TKQPX0KNBwm7Lovlpxjd_P_q4JsKWYrT3W_NKPr8w6oU7K8cYM-7bGuM9PFhNQEuBH109Z8yatDP0qApUVWMROZyFmBkFo3NEHBv0PZUm6pbxQU&u_asig=05DeK0a0DKFecR--vOFG3MvmtmBp9jxLSce_L0QO4yF1dTh-U3PJYrQ_gC3S3lWkDbqP0VU9cBI61wy3FHh7V0PTubf_IK4hrCLz3whd77hKVA21ELJhPXdYZN2uOZi1pTyIq5cgcUoW_xm5mG43bLn66504DH-y9yYoEQYDliY3T9JS7q8ZD7Xtz2Ly-b0kmuyAKRFSVJkkdwVUnyHAIJzcNJ8g9CawTHkVDn--N54Vh1A20HKBEoLwl1_A-SB6eWCs-hAuami3wC_ze-3aLJqu3h9VXwMyh6PgyDIVSG1W_DrNPko8qKmrDEL2lv2wjxnnfigo6gvybXhopmdXmQQM1fatBWw8lIqyLFhPoJ9kpcCjSnFubJWFWMMQTTneB_mWspDxyAEEo4kbsryBKb9Q&u_aref=BDjD4PSPlLGzPAkVuOsWAIORfW4%3D)

[go groupcache 用法示例_乃不知有汉的博客-CSDN博客](https://blog.csdn.net/wangjunsheng/article/details/81512767)



## 相关链接

[Java开发利器Guava Cache之使用篇 - 简书 (jianshu.com)](https://www.jianshu.com/p/7733c21cad7b?u_atoken=b91087f9-4283-4a76-bd4e-7bdcb9ad7fd1&u_asession=017RaHGAoRvhQLWSBFA4jSYpQY5KFXstvs-M8FVxZVDFntE7m50Igl8nl0qY9LaGeWX0KNBwm7Lovlpxjd_P_q4JsKWYrT3W_NKPr8w6oU7K9MW5b_f5s3SMbMa8BiYMz_9Z8yatDP0qApUVWMROZyFmBkFo3NEHBv0PZUm6pbxQU&u_asig=05oQsxl1GVuo1JWKDnpb89Yx5EPatFlMWszAr4KN1uVVoBT8XWRfj15Dgc_NXEDkGh5y4sxTftCHD_3bjwivgb61EV9edkCaBc6Mr1NDaFHJI8Gcqyy8rNn00vCbxFPARIdRix94en97SH4yvemZG32SKMpbLkQl1stNfOTCnZtZX9JS7q8ZD7Xtz2Ly-b0kmuyAKRFSVJkkdwVUnyHAIJzRkhgN2adPLMg8jMy62PUMrN2tL-EXqHEv-Le-Xp13d6Cs-hAuami3wC_ze-3aLJqu3h9VXwMyh6PgyDIVSG1W_5stPz4TcVwksKc-_gTlZ14SCiwcadFJVvu84nbKDsY0nDcJFezVD552H4jmxq1o8RpRsNCQx3W6_1ZVH_eCbTmWspDxyAEEo4kbsryBKb9Q&u_aref=vmEq9Vs5dLcoGhtCHNoSDqEypzw%3D)



