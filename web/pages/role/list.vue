<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addUser" type="primary">新增角色</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="ID" min-width="150" prop="roleId"></el-table-column>
      <el-table-column label="角色名" min-width="150" prop="role_name"></el-table-column>
      <el-table-column label="操作" min-width="150">
        <template slot-scope="scope">
          <el-button type="text" size="small" @click="editUser(scope.row)">编辑</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除此角色吗</p>
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
      <el-form :rules="formRules" ref="userForm" :model="formData">
        <el-form-item label="角色名" label-width="80px" prop="role_name">
          <el-input v-model="formData.role_name"></el-input>
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
import { nameValid, passwordValid } from "@/utils/validate.js";
import infoList from "@/components/mixins/infoList";
export default {
  name: "List",
  mixins: [infoList],
  data() {
    return {
      listApi: function() {
        return this.$axios.get("/role/list", {
          params: { page: this.page, pageSize: this.pageSize }
        });
      },
      isDialogShow: false,
      dialogTitle: "新增角色",
      formData: {
        role_name: "",
      },
      formRules: {
        role_name: nameValid,
      }
    };
  },
  async created() {
    this.getTableData();
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
      this.formData = row;
      this.openDialog("edit");
    },
    async deleteUser(row) {
      const { data: res } = await this.$axios.post("/role/delete", {
        roleId: row.roleId
      });
      if (res.code == 200) {
        this.getTableData();
        row.visible = false;
      }
    },
    async enterDialog() {
      this.$refs.userForm.validate(async valid => {
        if (valid) {
          let url = "role/" + this.actionType;
          const { data: res } = await this.$axios.post(url, this.formData);
          if (res.code !== 200) return this.$message.error(res.msg);
          this.$message.success(res.msg);
          await this.getTableData();
          this.closeDialog();
          this.initForm();
        }
      });
    },
    closeDialog() {
      this.$refs.userForm.resetFields();
      this.isDialogShow = false;
      this.initForm();
    },
    // 初始化表单
    initForm() {
      if (this.$refs.formData) {
        this.$refs.formData.resetFields();
      }
      this.formData = {
        role_name: "",
      };
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