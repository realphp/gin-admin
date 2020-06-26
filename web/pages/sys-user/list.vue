<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addUser" type="primary">新增用户</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="ID" min-width="100" prop="Id"></el-table-column>
      <el-table-column label="用户名" min-width="150" prop="user_name"></el-table-column>
      <el-table-column label="昵称" min-width="150" prop="nick_name"></el-table-column>
      <el-table-column label="角色" min-width="150" prop="role_id" :formatter="roleFormat"></el-table-column>
      <el-table-column label="联系方式" min-width="150" prop="phone"></el-table-column>
      <el-table-column label="操作" min-width="150">
        <template slot-scope="scope">
          <el-button type="text" size="small" @click="editUser(scope.row)">编辑</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除此用户吗</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteUser(scope.row)">确定</el-button>
            </div>
            <el-button type="text" size="small" slot="reference">删除</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :visible.sync="isDialogShow" custom-class="user-dialog" :title="dialogTitle">
      <el-form :rules="rules" ref="userForm" :model="userInfo">
        <el-form-item label="用户名" label-width="80px" prop="user_name">
          <el-input v-model="userInfo.user_name"></el-input>
        </el-form-item>
        <el-form-item label="昵称" label-width="80px" prop="nick_name">
          <el-input v-model="userInfo.nick_name"></el-input>
        </el-form-item>
        <el-form-item label="联系方式" label-width="80px" prop="phone">
          <el-input v-model="userInfo.phone"></el-input>
        </el-form-item>
        <el-form-item label="角色" label-width="80px" prop="role">
          <template>
            <el-select v-model="userInfo.role_id" placeholder="请选择">
              <el-option
                v-for="item in roles"
                :key="item.roleId"
                :label="item.role_name"
                :value="item.roleId"
              ></el-option>
            </el-select>
          </template>
        </el-form-item>
      </el-form>

      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API;
import { nameValid, mobileValid } from "@/utils/validate.js";
import infoList from "@/components/mixins/infoList";
export default {
  name: "List",
  mixins: [infoList],
  data() {
    return {
      listApi: function() {
        return this.$axios.get("/user/list", {
          params: { page: this.page, pageSize: this.pageSize }
        });
      },
      isDialogShow: false,
      dialogTitle: "新增用户",
      userInfo: {
        user_name: "",
        phone: "",
        nick_name: "",
        role_id: ""
      },
      roles: [],
      rules: {
        user_name: nameValid,
        phone: mobileValid,
        nick_name: nameValid
      }
    };
  },
  async created() {
    this.getTableData();
    const { data: res } = await this.$axios.get("/role/list", {
      params: { page: this.page, pageSize: this.pageSize }
    });
    this.roles = res.data.list;
    console.log(res);
  },
  methods: {
    openDialog(type) {
      switch (type) {
        case "add":
          this.dialogTitle = "新增";
          break;
        case "edit":
          this.dialogTitle = "编辑";
          break;
        default:
          break;
      }
      this.actionType = type;
      this.isDialogShow = true;
    },
    addUser() {
      this.openDialog("add");
    },
    editUser(row) {
      this.userInfo = row;
      this.openDialog("edit");
    },
    async deleteUser(row) {
      const { data: res } = await this.$axios.post("/user/delete", {
        id: row.Id
      });
      if (res.code == 200) {
        this.getTableData();
        row.visible = false;
      }
    },
    // 初始化表单
    initForm() {
      if (this.$refs.userForm) {
        this.$refs.userForm.resetFields();
      }
      this.userInfo = {
        user_name: "",
        phone: "",
        nick_name: "",
        role_id: ""
      };
    },
    // 确定弹窗
    async enterDialog() {
      this.$refs.userForm.validate(async valid => {
        if (valid) {
          let url = "user/" + this.actionType;
          const { data: res } = await this.$axios.post(url, this.userInfo);
          if (res.code !== 200) return this.$message.error(res.msg);
          this.$message.success(res.msg);
          await this.getTableData();
          this.closeDialog();
          this.initForm();
        }
      });
    },
    // 关闭窗口
    closeDialog() {
      this.$refs.userForm.resetFields();
      this.initForm();
      this.isDialogShow = false;
    },
    roleFormat(row) {
      for (let i = 0; i < this.roles.length; i++) {
        if (row.role_id == this.roles[i].roleId) {
          return this.roles[i].role_name;
        }
      }
    }
  }
};
</script>
<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}

.user-dialog {
  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }
  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}
</style>