<template>
  <div>
    <el-scrollbar style="height:calc(100vh)">
      <el-menu v-if="asyncRouters"
        class="el-menu-vertical"
        :collapse="isCollapse"
        :collapse-transition="true"
        :default-active="active"
        @select="selectMenuItem"
        active-text-color="#fff"
        text-color="rgb(191, 203, 217)"
        unique-opened
      >
        <template v-for="item in asyncRouters[0].children">
          <aside-component :key="item.name" :routerInfo="item" v-if="!item.hidden" />
        </template>
      </el-menu>
    </el-scrollbar>
  </div>
</template>

<script>
import { mapGetters, mapMutations } from "vuex";
import AsideComponent from "~/components/aside/asideComponent";
export default {
  name: "Aside",
  data() {
    return {
      active: "",
      isCollapse: false
    };
  },
  async fetch(){
    await this.$store.dispatch('router/SetAsyncRouter')
  },
  methods: {
    ...mapMutations("history", ["addHistory"]),
    selectMenuItem(index) {
      if (index === this.$route.name) return;
      this.$router.push(index);
    }
  },
  computed: {
    ...mapGetters("router", ["asyncRouters"])
  },
  components: {
    AsideComponent
  },
  created() {
    this.active = this.$route.name;
    if (process.browser) {
      let screenWidth = document.body.clientWidth;
      if (screenWidth < 1000) {
        this.isCollapse = !this.isCollapse;
      }

      this.$bus.on("collapse", item => {
        this.isCollapse = item;
      });
    }
  },
  watch: {
    $route() {
      this.active = this.$route.name;
    }
  },
  beforeDestroy() {
    this.$bus.off("collapse");
  }
};
</script>

<style lang="scss">
.el-scrollbar {
  .el-scrollbar__view {
    height: 100%;
  }
}
.menu-info {
  .menu-contorl {
    line-height: 52px;
    font-size: 20px;
    display: table-cell;
    vertical-align: middle;
  }
}
</style>