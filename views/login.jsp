<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>登录我找找</title>
    <meta name="viewport" content="width=device-width, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no" />
    <!-- <meta name="format-detection" content="telephone=no">
    <meta name="format-detection" content="email=no"> -->
    <link rel="stylesheet" href="css/style.css">
    <link rel="stylesheet" href="css/home/login.css">
</head>
<body>
<div class="wrapper">
    <section style="margin-top:70px;"></section>
    <section>
        <form action="login.action" method="post" onsubmit="return checkInput();">
            <ul>
                <li>
                    <input id="username" name="username" type="text" class="username" placeholder="请输入登录帐号">
                </li>
                <li>
                    <input id="password" name="password" type="password" class="pwd" placeholder="请输入密码">
                    <img src="images/eye.png" alt="eye.png" class="float-right" width="30px">
                </li>
            </ul>
            <div class="margin-bottom-25">
                <a href="registe.action" class="color-orange">找回密码</a>
                <div class="float-right">
                    <input id="autoLogin" name="autoLogin" type="checkbox">
                    <label for="autoLogin">自动登录</label>
                </div>
            </div>
            <button class="btn margin-bottom-25">登录</button>
            <div class="text-align-center">
                还没有帐号？
                <a href="registe.action" class="color-orange">立即注册</a>
            </div>
        </form>
    </section>
</div>

<script type="text/javascript" src="js/jquery-2.1.4.min.js"></script>
<script type="text/javascript" src="js/Validator.js"></script>
<script type="text/javascript">
    function checkInput() {
        var username = $('#username').val();
        if (Validator.isNull(username)) {
            alert('请输入用户名');
            return false;
        }
        if (Validator.isNotTelAndPhone(username)) {
            alert('请输入手机号码进行登录');
            return false;
        }
        var password = $('#password').val();
        if (Validator.isNull(password)) {
            alert('请输入密码');
            return false;
        }
        return true;
    }
</script>
<%@ include file="../tips.jsp"%>
</body>
</html>