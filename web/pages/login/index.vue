<template>
  <div id="userLayout" class="user-layout-wrapper">
    <div class="container">
      <div class="top">
        <div class="header">
          <a href="/">
            <!-- <img src="~@/assets/logo.png" class="logo" alt="logo" /> -->
            <span class="title">Gin-Admin</span>
          </a>
        </div>
      </div>
      <div class="main">
        <el-form
          :model="loginForm"
          :rules="loginFormRules"
          ref="loginForm"
          @keyup.enter.native="submitForm"
        >
          <el-form-item prop="username">
            <el-input placeholder="请输入用户名" v-model="loginForm.user_name">
              <i class="el-input__icon el-icon-user" slot="suffix"></i>
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input type="password" placeholder="请输入密码" v-model="loginForm.password">
              <i class="el-input__icon el-icon-lock"  slot="suffix"></i>
            </el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm" style="width:100%">登 录</el-button>
          </el-form-item>
        </el-form>
      </div>

      <div class="footer">
        <div class="links"></div>
        <div class="copyright">Copyright &copy; 💖flipped-aurora</div>
      </div>
    </div>
  </div>
</template>
<script>
import { nameValid, passwordValid } from "@/utils/validate.js";
export default {
  name: "Login",
  layout:"blank",
  data() {
    return {
      // 表单数据绑定对象
      loginForm: {
        user_name: "",
        password: ""
      },
      // 表单的规则验证对象
      loginFormRules: {
        user_name: nameValid,
        password: passwordValid
      }
    };
  },
  methods: {
    async submitForm() { 
      this.$refs.loginForm.validate(async valid => {
        if (!valid) return;
        const { data: res } = await this.$axios.post("login", this.loginForm);
        if (res.code !== 200) return this.$message.error(res.msg);
        this.$message.success("登录成功");
        // 保存到客户端的Cookie中
        this.$cookies.set("token", res.data.token);
        // 保存用户信息到vuex
        this.$store.dispatch("setUser", res);
        this.$router.push("/");
      });
    }
  }
};
</script>

<style scoped lang="scss">
.login-register-box {
  height: 100vh;
  .login-box {
    width: 40vw;
    position: absolute;
    left: 50%;
    margin-left: -22vw;
    top: 5vh;
    .logo {
      height: 35vh;
      width: 35vh;
    }
  }
}
.link-icon {
  width: 20px;
  min-width: 20px;
  height: 20px;
  border-radius: 10px;
}
.vPic {
  width: 33%;
  height: 38px;
  float: right !important;
  background: #ccc;
  img {
    cursor: pointer;
    vertical-align: middle;
  }
}
.logo_login {
  width: 100px;
}
#userLayout.user-layout-wrapper {
  height: 100%;
  position: relative;
  &.mobile {
    .container {
      .main {
        max-width: 368px;
        width: 98%;
      }
    }
  }
  .container {
    position: relative;
    overflow: auto;
    width: 100%;
    min-height: 100%;
    background-size: 100%;
    padding: 110px 0 144px;
    a {
      text-decoration: none;
    }
    .top {
      text-align: center;
      margin-top: -40px;
      .header {
        height: 44px;
        line-height: 44px;
        margin-bottom: 30px;
        .badge {
          position: absolute;
          display: inline-block;
          line-height: 1;
          vertical-align: middle;
          margin-left: -12px;
          margin-top: -10px;
          opacity: 0.8;
        }
        .logo {
          height: 44px;
          vertical-align: top;
          margin-right: 16px;
          border-style: none;
        }
        .title {
          font-size: 33px;
          color: rgba(0, 0, 0, 0.85);
          font-family: Avenir, "Helvetica Neue", Arial, Helvetica, sans-serif;
          font-weight: 600;
          position: relative;
          top: 2px;
        }
      }
      .desc {
        font-size: 14px;
        color: rgba(0, 0, 0, 0.45);
        margin-top: 12px;
      }
    }
    .main {
      min-width: 260px;
      width: 368px;
      margin: 0 auto;
    }
    .footer {
      position: relative;
      width: 100%;
      margin: 40px 0 10px;
      text-align: center;
      .links {
        margin-bottom: 8px;
        font-size: 14px;
        width: 330px;
        display: inline-flex;
        flex-direction: row;
        justify-content: space-between;
        padding-right: 40px;
        a {
          color: rgba(0, 0, 0, 0.45);
          transition: all 0.3s;
        }
      }
      .copyright {
        color: rgba(0, 0, 0, 0.45);
        font-size: 14px;
        padding-right: 40px;
      }
    }
  }
}
</style>
