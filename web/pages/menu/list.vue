<template>
  <div>
    <el-row>
      <el-button type="primary" icon="el-icon-plus" size="mini" @click="handleAdd">新增</el-button>
    </el-row>

    <el-table :data="menus" row-key="ID" border stripe :tree-props="{children: 'children'}">
      <el-table-column prop="meta.title" label="菜单名称" :show-overflow-tooltip="true" width="180px" />
      <el-table-column prop="meta.icon" label="图标" align="center" width="100px">
        <template slot-scope="scope">
          <i :class="'el-icon-'+scope.row.meta.icon"></i>
        </template>
      </el-table-column>
      <el-table-column prop="sort" label="排序" width="60px" />
      <el-table-column prop="component" label="路径" :show-overflow-tooltip="true"></el-table-column>
      <el-table-column prop="hidden" label="可见" :formatter="visibleFormat" width="80">
        <template slot-scope="scope">
          <el-tag
            :type="scope.row.hidden === 'true' ? 'danger' : 'success'"
            disable-transitions
          >{{ visibleFormat(scope.row) }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" align="center" prop="createdAt" width="180">
        <template slot-scope="scope">
          <span>{{ scope.row.CreatedAt|formatDate}}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width" width="180">
        <template slot-scope="scope">
          <el-button size="mini" type="text" icon="el-icon-edit" @click="handleUpdate(scope.row)">修改</el-button>
          <el-button
            size="mini"
            type="text"
            icon="el-icon-delete"
            @click="handleDelete(scope.row)"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
 
<script>
import { formatTimeToStr } from "@/utils/date";
export default {
  data() {
    return {
      menus: []
    };
  },
  async created() {
    const { data: res } = await this.$axios.post("/menu/list");
    this.menus = res.Menus;
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    }
  },
  methods: {
    // 菜单显示状态字典翻译
    visibleFormat(row) {
      if (row.hidden === "false") {
        return "隐藏";
      }
      return "显示";
    },
    handleUpdate() {},
    handleAdd() {},
    handleDelete() {}
  }
};
</script> 
<style>
</style>