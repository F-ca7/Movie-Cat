<!-- 头部区 -->
<header class="header headerStyle1" id="header">
    <div class="sticky scrollHeaderWrapper">
        <div class="container">
            <div class="row">

                <!-- logo栏-->
                <div class="logoWrapper">
                    <h1 class="logo">
                        <a class="clearfix" href="/index">
                    <span class="square">
                      <span>F</span>
                    </span>
                            <span class="text">Movie cat</span>
                        </a>
                    </h1><!-- end of logo -->
                </div><!-- end of logoWrapper -->

                <!-- 导航栏-->
                <nav class="mainMenu mainNav" id="mainNav">
                    <ul class="navTabs">
                        <li>
                            <a>首页</a>
                            <ul class="dropDown sub-menu">
                                <li><a href="/index#current-show">当前热映</a></li>
                                <li><a href="/index#recommend">推荐观影</a></li>
                                <li><a href="/index#domainSearch">搜索电影</a></li>
                                <li><a href="/index#remark">精选影评</a></li>
                            </ul><!-- end of 子菜单 -->
                        </li>
                        <li>
                            <a>类别筛选</a>
                            <ul class="dropDown sub-menu">
                                <li><a href="/list?category=剧情">剧情</a></li>
                                <li><a href="/list?category=爱情">爱情</a></li>
                                <li><a href="/list?category=犯罪">犯罪</a></li>
                                <li><a href="/list?category=惊悚">惊悚</a></li>
                                <li><a href="/list?category=冒险">冒险</a></li>
                                <li><a href="/list?category=科幻">科幻</a></li>
                                <li><a href="/list?category=动作">动作</a></li>
                                <li><a href="/list?category=喜剧">喜剧</a></li>
                                <li><a href="/list?category=奇幻">奇幻</a></li>
                                <li><a href="/list?category=悬疑">悬疑</a></li>
                                <li><a href="/index#domainSearch">关键词搜索</a></li>

                            </ul><!-- end of 子菜单 -->
                        </li>
                        <li>
                            <a>排行榜</a>
                            <ul class="dropDown sub-menu">
                                <li>
                                    <a href="/top250">TOP250排行榜</a>
                                </li>
                                <li>
                                    <a href="/top250">IMDB排行榜</a>
                                </li>
                            </ul><!-- end of 子菜单 -->
                        </li>
                        <!-- 似乎没时间做影评 -->
                        {{/*<li>*/}}
                            {{/*<a>影评</a>*/}}
                            {{/*<ul class="dropDown sub-menu">*/}}
                                {{/*<li>*/}}
                                    {{/*<a href="/list">豆瓣影评</a>*/}}
                                {{/*</li>*/}}
                                {{/*<li>*/}}
                                    {{/*<a href="/list">IMDB影评</a>*/}}
                                {{/*</li>*/}}
                            {{/*</ul><!-- end of 子菜单 -->*/}}
                        {{/*</li>*/}}

                        <!-- 登录后台逻辑有时间开发-->
                        <li class="login formTop">
                            {{if .IsLogin}}
                                <button class="formSwitcher"  data-toggle="modal" data-target="#loginModal"><a href="/logout">登出</a></button>
                            {{else}}
                            <button class="formSwitcher"  data-toggle="modal" data-target="#loginModal">登录</button>
                            <div class="modal loginModal fade" id="loginModal" tabindex="-1" role="dialog" aria-hidden="true">
                                <div class="container">
                                    <ol class="formWrapper loginFormWrapper" id="loginFormWrapper">
                                        <li><h5><i class="fa fa-user"></i>登录Movie Cat</h5></li>
                                        <li>
                                            <form class="loginForm" method="POST" action="/login">
                                                <input class="loginName" id="loginName" name="username" placeholder="用户名" type="text" >
                                                <input class="loginPassword" id="loginPassword" name="password" placeholder="密码" type="password">
                                                <button class="generalBtn loginBtn" type="submit">登录</button>
                                            </form>
                                        </li>
                                        <li class="register"><p><a href="/register">创建新账户</a></p></li>
                                    </ol>
                                </div><!-- end of container -->
                            </div><!-- end of 登录框 -->
                            <a href="login.html">登录</a>
                            {{end}}
                        </li>
                    </ul><!-- end of 导航项 -->
                </nav><!-- end of 导航栏 -->

                <a href="#" class="generalLink" id="responsiveMainNavToggler"><i class="fa fa-bars"></i></a>
                <div class="clearfix"></div><!-- 清除浮动-->
                <div class="responsiveMainNav"></div><!-- end of  -->

            </div><!-- end of 当前行 -->
        </div><!-- end of container -->
    </div>
</header><!-- end of 头部 -->