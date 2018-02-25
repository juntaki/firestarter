<template>
<div>
  <el-button type="text" @click="showDeleteDialog=true">Delete</el-button>
  <el-dialog :title="'Delete ' + config.title" width="400px" :visible.sync="showDeleteDialog">
    <span>Are you sure to delete this config?</span>
    <span slot="footer" class="dialog-footer">
      <el-button type="danger">Yes</el-button><el-button>Cancel</el-button>
    </span>
  </el-dialog>
  </div>
</template>

<script>
export default {
  props: ['config'],
  data () {
    return {
      showDeleteDialog: false
    }
  },
  methods: {
    deleteConfig () {
    }
  }
}

</script>

<style scoped>

</style>
