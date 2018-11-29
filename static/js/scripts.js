/*
   自定义js效果
*/

(function ()
{
  //加载动画
  jQuery(window).load(function(){
    jQuery('.loadingContainer').css({'opacity' : 0 , 'display' : 'none'});
    jQuery('.allWrapper').css({'opacity' : 1 , 'visibility' : 'visible'});
    jQuery('.switcher').css({'opacity' : 1 , 'visibility' : 'visible'});
    jQuery('.back-to-top').css({'opacity' : 1 , 'visibility' : 'visible'});
    jQuery('body').css({'overflow' : 'visible'});
  });



  //根据窗口大小调节元素尺寸
  var jQuerywindow = jQuery(window);
  
  jQuerywindow.resize(function(){
    //固定item高度
    jQuery('#slider  , .slider .item > img , #banner , #banner .item').css({ 'height' : "783px" });
  });

  jQuerywindow.trigger('resize');

  //自定义container位置
  var jQueryloadingContainer = jQuery('.loadingContainer');

  jQueryloadingContainer.resize(function(){
    jQuery('.loadingContainer').css({ 
      'margin-left' : - jQueryloadingContainer.width() / 2 , 
      'margin-top' : - jQueryloadingContainer.height()  / 2
    });
  });

  jQueryloadingContainer.trigger('resize');


  //固定header区
  var jQueryheader = jQuery('header#header');
  var jQueryheaderTop = jQueryheader.offset().top;
  jQuery('.offset').height( jQueryheader.outerHeight() )
  jQuery(window).scroll(function(){
    (jQuery(window).scrollTop() > jQueryheaderTop) ? jQuery('.header').addClass('fixedHeader') : jQuery('.header').removeClass('fixedHeader');
  });





   //监听页面的滚动事件
  jQuery('.scrollTo').on('click', function( e ) {

      var scrollAnchor = jQuery(this).attr('data-scroll'),
          scrollPoint = jQuery('.scrollAnchor[data-anchor="' + scrollAnchor + '"]').offset().top - 28;

      jQuery('body,html').animate({
          scrollTop: scrollPoint
      }, 1500);

      window.location.hash = jQuery(this).attr('href').replace('#','');
      e.preventDefault();
  })


  //主页海报轮播
  jQuery(document).ready(function() {

    jQuery(".homeSlider_1 , .postSlider").owlCarousel({
      animateOut: 'fadeOut',
      items:1,
      margin:0,
      loop:true,
      autoplay:true,
      autoplayTimeout:8000,
      autoplayHoverPause:false,
      nav: true,
      dots: false,
      stagePadding:0,
      smartSpeed:1000,
      responsive:{
        0:{
          items:1
        },
        768:{
          items:1
        },
        1000:{
          items:1
        }
      }
    });

  });

  //主页轮播2
  jQuery(document).ready(function() {

    var owl = jQuery(".homeSlider_2");

    owl.owlCarousel({
      animateOut: 'fadeOut',
      items:1,
      margin:0,
      loop:true,
      autoplay:true,
      autoplayTimeout:8000,
      autoplayHoverPause:false,
      nav: true,
      dots : true,
      responsive:{
          0:{
              items:1
          },
          768:{
              items:1
          },
          1000:{
              items:1
          }
        }
    });
  });




  //轮播1
  jQuery(document).ready(function() {

    var owl = jQuery(".carousel");

    owl.owlCarousel({
      nav : false,
      dots : true,
      loop:true,
      autoplay: true,
      items:4,
      responsive:{
        0:{
            items:1,
            slideBy: 1
        },
        768:{
            items:3,
            slideBy: 3
        },
        1000:{
            items:4,
            slideBy: 4
        }
      }
    });
  });



    //轮播2
  jQuery(document).ready(function() {

        var owl = jQuery(".carousel2");

        owl.owlCarousel({
            nav : true,
            dots : false,
            animateOut : 'slideUpSlow',
            animateIn: 'fadeOutUp',
            autoplay:true,
            loop:true,
            items : 1,
            responsive:{
                0:{
                    items:1
                },
                768:{
                    items:1
                },
                1000:{
                    items:1
                }
            }
        });
    });

    //轮播3
  jQuery(document).ready(function() {

        var owl = jQuery(".carousel3");

        owl.owlCarousel({
            nav : true,
            dots : false,
            animateOut : 'slideUpSlow',
            animateIn: 'fadeOutUp',
            autoplay:true,
            loop:true,
            items : 1,
            responsive:{
                0:{
                    items:1
                },
                768:{
                    items:1
                },
                1000:{
                    items:1
                }
            }
        });
    });

  //精选评论区轮播
  jQuery(document).ready(function() {

    var owl = jQuery(".testmonialsCarousel");

    owl.owlCarousel({
      nav : false,
      dots : true,
      loop:true,
      autoplay : true,
      items : 3,
      responsive:{
        0:{
            items:1,
            slideBy: 1
        },
        768:{
            items:2,
            slideBy: 2
        },
        1000:{
            items:3,
            slideBy: 3
        }
      }
    });
  });


  //影评页面轮播
  jQuery(document).ready(function() {

    var owl = jQuery(".remarkCarousel2");

    owl.owlCarousel({
      nav : true,
      dots : false,
      animateOut : 'fadeOut',
      animateIn: 'flipInX',
      loop:true,
      autoplay : true,
      items : 1,
      responsive:{
        0:{
            items:1
        },
        768:{
            items:1
        },
        1000:{
            items:1
        }
      }
    });
  });



  /* 控制输入框的淡入淡出 */
  //影评输入
  jQuery(document).ready(function() {
    
    var watermark = 'Comment';
    
    //init, set watermark text and class
    jQuery('#commentArea').val(watermark).addClass('inputBar');
    
    //if blur and no value inside, set watermark text and class again.
    jQuery('#commentArea').blur(function(){
        if (jQuery(this).val().length == 0){
          jQuery(this).val(watermark).addClass('inputBar');
      }
    });

    //if focus and text is watermrk, set it to empty and remove the watermark class
    jQuery('#commentArea').focus(function(){
        if (jQuery(this).val() == watermark){
          jQuery(this).val('').removeClass('inputBar');
      }
    });
  });


  //回到顶部
  jQuery(document).ready(function() {
    var offset = 100;       //todo 修改按钮出现的时机
    var duration = 1000;    //todo 修改滑动速度
    jQuery(window).scroll(function() {
        if (jQuery(this).scrollTop() > offset) {
            jQuery('.back-to-top').addClass('fadeInup');      //显示按钮
        } else {
            jQuery('.back-to-top').removeClass('fadeInup');   //隐藏按钮
        }
    });
    
    jQuery('.back-to-top').click(function(e) {
        e.stopPropagation();
        jQuery('body,html').animate({
        scrollTop: 0
    }, duration);
        return false;
    })
  });


  //返回顶部按钮
  jQuery(document).ready(function() {
      jQuery('body').append(
          "<a href='#' class='back-to-top'>"+
          "<i class='fa fa-chevron-up'></i>"+
          "</a>"
      );
  });




  /* 滑动效果 */

  //页面内效果
  jQuery("html").niceScroll({
    cursorborder: 0,
    cursorcolor: '#171717',
    autohidemode: true,
    zindex: 9999999,
    scrollspeed: 60,
    mousescrollstep: 36,
    cursorwidth: 6,
    horizrailenabled: false,
    cursorborderradius: 3
  });


  //模态框插件
  jQuery(".modal").niceScroll({
    cursorborder: 0,
    cursorcolor: '#fff',
    autohidemode: true,
    zindex: 99999999999,
    scrollspeed: 60,
    mousescrollstep: 36,
    cursorwidth: 6,
    horizrailenabled: false,
    cursorborderradius: 3
  });



  /* 导航栏的响应与固定 */
 
  jQuery(document).ready(function () {

    jQuery('.mainNav').clone().appendTo('.responsiveMainNav');

    //控制隐藏与显示导航栏
    jQuery('#responsiveMainNavToggler').click(function(event){
      event.preventDefault();
      jQuery('#responsiveMainNavToggler').toggleClass('opened');
      jQuery('.responsiveMainNav').slideToggle(1000);


      if ( jQuery('#responsiveMainNavToggler i').hasClass('fa-bars') )
      {   
          jQuery('#responsiveMainNavToggler i').removeClass('fa-bars');
          jQuery('#responsiveMainNavToggler i').addClass('fa-close');
      }else
      {  
          jQuery('#responsiveMainNavToggler i').removeClass('fa-close');
          jQuery('#responsiveMainNavToggler i').addClass('fa-bars');
      }

    });

    // 一级下拉菜单（要加再加，先不考虑扩展性，太麻烦）
    if(jQuery(".responsiveMainNav .navTabs > li > a").parent().has("ul")) {
      jQuery(".responsiveMainNav .navTabs > li > a:first-child").addClass("toggleResponsive");
      jQuery(".responsiveMainNav .navTabs > li > a:last-child").removeClass("toggleResponsive");
    }

    jQuery(".responsiveMainNav .navTabs > li > .toggleResponsive").on("click", function(e){
      if(jQuery(this).parent().has("ul")) {
        e.preventDefault();
      }
      
      if(!jQuery(this).hasClass("activeLine")) {
        // hide any open menus and remove all other classes
        jQuery(".responsiveMainNav .navTabs > li > .toggleResponsive").removeClass("activeLine");
        jQuery(".responsiveMainNav .navTabs > li > .dropDown").slideUp(500);
        
        // open our new menu and add the activeLine class
        jQuery(this).addClass("activeLine");
        jQuery(this).next(".responsiveMainNav .navTabs > li > .dropDown").slideDown(500);
      }
      
      else if(jQuery(this).hasClass("activeLine")) {
        jQuery(this).removeClass("activeLine");
        jQuery(this).next(".responsiveMainNav .navTabs > li > .dropDown").slideUp(500);
      }
    });


    jQuery(".responsiveMainNav .navTabs > li > .dropDown > li > .toggleResponsive").on("click", function(e){
      if(jQuery(this).parent().has("ul")) {
        e.preventDefault();
      }

      if(!jQuery(this).hasClass("activeLine")) {
        // hide any open menus and remove all other classes
        jQuery(".responsiveMainNav .navTabs > li > .dropDown > li > .toggleResponsive").removeClass("activeLine");
        jQuery(".responsiveMainNav .navTabs > li > .dropDown li .dropDown").slideUp(500);
        
        // open our new menu and add the activeLine class
        jQuery(this).addClass("activeLine");
        jQuery(this).next(".responsiveMainNav .navTabs > li > .dropDown li .dropDown").slideDown(500);
      }
      
      else if(jQuery(this).hasClass("activeLine")) {
        jQuery(this).removeClass("activeLine");
        jQuery(this).next(".responsiveMainNav .navTabs > li > .dropDown li .dropDown").slideUp(500);
      }
    });

  });

  
})();

$(function(){
    var len=$(".main_banner li").length;
    var index_2=0;
    var timer=800;
    var intervaltimer=0;
    var isMoving=false;

    function slide(slideMode){//轮播方法
        if (isMoving==false){
            isMoving=true;
            var prev; var next; var hidden;
            var curr=$("#imgCard"+index_2);//当前正中显示

            if(index_2==0){								//当前正中显示的是第0张时 prev为最后一张
                prev=$("#imgCard"+(len-1));
            }else{												//否则  序列号-1
                prev=$("#imgCard"+(index_2-1));
            }
            if(index_2==(len-1)){					//当前正中显示的是最后一张时 next为第0张
                next=$("#imgCard0");
            }else{											//否则  序列号+1
                next=$("#imgCard"+(index_2+1));
            }

            if(slideMode){			//slideMode为1(true)，执行slide(1)，上一张
                if(index_2-2>=0){									//index_2						2		3		4
                    hidden=$("#imgCard"+(index_2-2));//									0		1		2
                }else{													//index_2		0		1
                    hidden=$("#imgCard"+(len+index_2-2));//			3		4
                }
                prev.css("z-index","5");			//点击prev按钮  让prev位置上的这张图片 层级最高 显示
                next.css("z-index","1");
                curr.css("z-index","2");
                hidden.css("z-index","1");
                //当index_2自减，各图片往右运动效果
                hidden.css({width:"450px",height:"220px",top:"60px","left":"0px","opacity":0});
                hidden.stop(true,true).animate({width:"580px",height:"280px",top:"20px",left:"0px",opacity:1},timer);
                curr.stop(true,true).animate({width:"580px",height:"280px",top:"20px",left:"600px",opacity:1},timer);
                next.stop(true,true).animate({width:"450px",height:"220px",top:"60px","left":"730px","opacity":0},timer,function(){next.find("span").css("opacity",0); isMoving = false;});
                //prev  -->  curr     prev中的图片li轮换到curr的位置      其他一次轮换
                prev.find("span").css("opacity",0);
                $(".main_banner_box li").find("p").css({"bottom":"-50px"});//所有标题隐藏
                prev.stop(true,true).animate({width:"670px",height:"320px",left:"255px",top:0,opacity:1},timer,function(){
                    $(this).find("p").animate({"bottom":"0px"});	//当前这张图片的标题运动出来
                });
                index_2--;
            }else{			//执行next 操作
                if(index_2+2>=len){								//index_2								3		4
                    hidden=$("#imgCard"+(index_2+2-len));//										0		1
                }else{													//index_2		0		1		2
                    hidden=$("#imgCard"+(index_2+2));//						2		3		4
                }
                prev.css("z-index","1");
                next.css("z-index","5");			//点击next按钮  让next位置上的这张图片 层级最高 显示
                curr.css("z-index","2");
                hidden.css("z-index","1");
                //当index_2自增，各图片往左运动效果
                hidden.css({width:"450px",height:"220px",top:"60px","left":"730px","opacity":0});
                hidden.stop(true,true).animate({width:"580px",height:"280px",top:"20px",left:"600px",opacity:1},timer);
                curr.stop(true,true).animate({width:"580px",height:"280px",top:"20px",left:"0px",opacity:1},timer);
                //next  -->  curr     next中的图片li轮换到curr的位置      其他一次轮换
                next.find("span").css("opacity",0);
                $(".main_banner_box li").find("p").css({"bottom":"-50px"});//所有标题隐藏
                next.stop(true,true).animate({width:"670px",height:"320px",left:"255px",top:0,opacity:1},timer,function(){
                    $(this).find("p").animate({"bottom":"0px"});	//当前这张图片的标题运动出来
                });
                prev.stop(true,true).animate({width:"450px",height:"220px",left:"0px",top:"60px",opacity:0},timer,function(){
                    isMoving = false;
                });
                index_2++;
            }//if else

            hidden.find("span").css("opacity",0.5);
            curr.find("span").css("opacity",0.5);

            if(index_2==len) index_2=0;
            if(index_2<0) index_2=len+index_2;			//限制index_2的范围
            $(".btn_list span").removeClass('curr').eq(index_2).addClass('curr');//给序列号按钮添加、移除样式
        }
    }//slide()


    if(len>3){
        //序列号按钮 跳序切换 方法
        $(".btn_list span").click(function(event){

            if (isMoving ) return;
            var oIndex=$(this).index();

            if(oIndex==index_2) return;//点击按钮的序列号与当前图片的序列号一致，return
            clearInterval(intervaltimer)
            intervaltimer=null;

            var flag=false;
            //当前显示图片的序列号  和  被点击按钮的序列号  间隔超过1且不是首尾两个的时候
            if(Math.abs(index_2-oIndex)>1&&Math.abs(len-Math.abs(index_2-oIndex))!=1){
                //统一样式
                $(".main_banner_box li").css({width:"300px",height:"120px",left:"600px",top:"60px",opacity:0});
                //如果当前的序列号   比    被点击按钮序列号     大     而且     不相邻、不是首尾
                if(index_2>oIndex&&len-Math.abs(index_2-oIndex)!=1){
                    flag=true;
                    index_2=oIndex+1;		//oIndex+1    通过slide()  运动回上一张    oIndex
                }else{//比   小     而且     不相邻、不是首尾
                    index_2=oIndex-1;		//oIndex-1     通过slide()  运动到下一张    oIndex
                    if(index_2<0) index_2=len-1;
                }
            }else{//当前 比 被点击  大	且   相邻									//从0    跳到     4		要执行上一张方法
                if((index_2>oIndex&&len-(index_2-oIndex)!=1)||(index_2<oIndex&&len+(index_2-oIndex)==1)){
                    flag=true;			//执行上一张
                }
            }
            slide(flag);
            intervaltimer=setInterval(slide,3000);//自动轮播

        });

        $(".main_banner_box li").on("mousemove",function(){
            if($(this).css("width")=="670px"){//鼠标移入为当前正中显示的图片li，则清除定时器
                clearInterval(intervaltimer);
                intervaltimer=null;
            }
        }).on("mouseout",function(){//鼠标移除重新滚动
            clearInterval(intervaltimer);
            intervaltimer=null;
            intervaltimer=setInterval(slide,3000);
        });

        $(".js_pre").click(function(event){//上一张
            if (isMoving ) return;
            clearInterval(intervaltimer);
            intervaltimer=null;
            slide(1);
            intervaltimer=setInterval(slide,3000);
        });

        $(".js_next").click(function(event){//下一张
            if (isMoving ) return;
            clearInterval(intervaltimer);
            intervaltimer=null;
            slide();
            intervaltimer=setInterval(slide,3000);
        });

        intervaltimer=setInterval(slide,3000);

    }else{

        $(".js_pre").hide();
        $(".js_next").hide();

    }//if else

});


    
