<template>
<div>
  <el-card>

    <el-button type="" @click="showDialog=true">Edit</el-button>
  </el-card>
  <el-dialog :visible.sync="showDialog">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="Channels">
        <el-select v-model="form.channels" placeholder="Channels"
          multiple auto-complete filterable style="width: 100%">
          <el-option v-for="item in channel" :key="item" :label="item" :value="item"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="Text Template">
        <el-input v-model="form.text" placeholder="Text Template"></el-input>
      </el-form-item>
      <el-form-item label="Actions">
        <el-select v-model="form.actions" placeholder="Actions"
          multiple allow-create filterable style="width: 100%"
          no-data-text="Please input action">
        </el-select>
      </el-form-item>
      <el-form-item label="URL Template">
        <el-input v-model="form.urlTemplate" placeholder="URL Template"></el-input>
      </el-form-item>
      <el-form-item label="Body Template">
        <el-input v-model="form.bodyTemplate" placeholder="Body Template"></el-input>
      </el-form-item>

      <el-form-item label="Confirm">
        <el-switch v-model="form.confirm"></el-switch>
      </el-form-item>

      <el-button type="primary" style="width: 100%">Submit</el-button>
    </el-form>
  </el-dialog>
  </div>
</template>

<script>
export default {
  data () {
    return {
      showDialog: false,
      channel: [
        'hoge', 'bottest', 'hoge1', 'bottest1'
      ],
      form: {
        channels: ['hoge'],
        actions: ['hoge'],
        text: 'sample',
        id: 'hoge',
        urlTemplate: 'u tmp',
        bodyTemplate: 'body',
        confirm: true
      }
    }
  },
  methods: {
    onSubmit () {
      console.log('submit!')
    }
  }
}

</script>

<style scoped>

</style>
