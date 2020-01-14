<template>
  <div class="container">
    <div>123</div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'LoginLog',
  computed: {
    ...mapGetters([
      // 'name'
    ])
  },
  // eslint-disable-next-line vue/order-in-components
  data() {
    return {

    }
  }
}
</script>

<style lang="scss" scoped>
.container {
  padding: 20px;
  .filter-container {
    padding-bottom: 10px;
    .filter-item {
      float: right;
    }
  }
}
.clearfix:after {
  /*伪元素是行内元素 正常浏览器清除浮动方法*/
  content: "";
  display: block;
  height: 0;
  clear: both;
  visibility: hidden;
}
.clearfix {
  *zoom: 1; /*ie6清除浮动的方式 *号只有IE6-IE7执行，其他浏览器不执行*/
}
</style>