<template>
<div>
  <template v-if="newConfig">
    <el-button style="width:100%" type="primary" @click="showDialog=true">Add config</el-button>
  </template>
  <template v-else>
    <el-button type="text" @click="showDialog=true">Edit</el-button>
  </template>
  <el-dialog :title="title()" width="90%" :visible.sync="showDialog">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="Title" prop="title" :rules="[{ required: true, message: 'Please input title', trigger: 'change' }]">
        <el-input v-model="form.title" placeholder="Deploy bot for my team"></el-input>
      </el-form-item>

      <h3>Trigger</h3>

      <el-form-item label="Channels" prop="channelsList"
        :rules="[{ required: true, message: 'Please input channels', trigger: 'change' }]">
        <el-select v-model="form.channelsList" placeholder="general random"
          multiple auto-complete filterable allow-create style="width: 100%">
          <el-option v-for="item in channels" :key="item" :label="item" :value="item"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="Regexp" prop="regexp"
      :rules="[{ required: true, message: 'Please input Regexp', trigger: 'change' }]">
        <el-input v-model="form.regexp" placeholder="^depoy (.*)$"></el-input>
      </el-form-item>

      <h3>Bot message</h3>

      <el-form-item label="Text" prop="text"
      :rules="[{ required: true, message: 'Please input Text', trigger: 'change' }]">
        <el-input v-model="form.text" placeholder="Deploy my app"></el-input>
      </el-form-item>
      <el-form-item label="Actions">
        <el-select v-model="form.actionsList" placeholder="master branch1 branch2"
          multiple allow-create filterable style="width: 100%"
          no-data-text="Please input action">
        </el-select>
      </el-form-item>
      <el-form-item label="Confirm">
        <el-switch v-model="form.confirm"></el-switch>
      </el-form-item>

      <h3>POST Request</h3>

      <el-form-item label="URL Template" prop="urltemplate"
      :rules="[{ required: true, message: 'Please input URL Template', trigger: 'change' }]">
        <el-input v-model="form.urltemplate" :placeholder="urlTemplatePlaceholder"></el-input>
      </el-form-item>
      <el-form-item label="Body Template">
        <el-input v-model="form.bodytemplate" :placeholder="bodyTemplatePlaceholder"></el-input>
      </el-form-item>

      <h3>Secrets</h3>

      <el-form-item
        v-for="(secret, index) in secrets"
        :label="'Secret (' + index + ')'"
        :key="secret.key"
        prop="secrets"
      >
        <el-row>
          <el-col :span="10"><el-input v-model="secret.secretKey" placeholder="AWS_TOKEN"></el-input></el-col>
          <el-col :span="10"><el-input v-model="secret.secretValue" placeholder="AKIAXXXXX"></el-input></el-col>
          <el-col :span="4"><el-button @click.prevent="removeSecret(secret)" style="width: 100%">Delete</el-button></el-col>
        </el-row>
      </el-form-item>
      <el-form-item>
        <el-button @click="addSecret">New secret</el-button>
      </el-form-item>

      <div><el-button type="primary" style="width: 100%" @click="onSubmit()">Submit</el-button></div>
      <div v-if="!newConfig">
        <el-button type="danger" style="width: 100%" @click="showDeleteDialog=true">Delete this config</el-button>
        <el-dialog :title="'Delete ' + config.title" width="400px" :visible.sync="showDeleteDialog" append-to-body="">
          <span>Are you sure to delete this config?</span>
          <span slot="footer" class="dialog-footer">
            <el-button type="danger" @click="deleteConfig()">Yes</el-button><el-button @click="showDeleteDialog=false">Cancel</el-button>
          </span>
        </el-dialog>
      </div>
    </el-form>
  </el-dialog>
  </div>
</template>

<script>
import config from '../../proto/config_pb_twirp'
export default {
  props: ['config', 'channels'],
  data () {
    const form = this.config
      ? JSON.parse(JSON.stringify(this.config))
      : {
        id: null,
        secrets: new Map()
      }

    const secrets = []
    form.secrets.forEach((v, k, m) => {
      secrets.push({
        key: secrets.length,
        secretKey: k,
        secretValue: v
      })
    })

    const host = location.protocol + '//' + location.host
    return {
      client: config.createConfigServiceClient(host),
      showDialog: false,
      showDeleteDialog: false,
      form: form,
      secrets: secrets,
      urlTemplatePlaceholder: 'https://example.com/deploy?param={{index .matched 1}}&value={{value}}',
      bodyTemplatePlaceholder: "{ value: '{{value}}' }"
    }
  },
  watch: {
    secrets: {
      handler: function (newValue, oldValue) {
        const newSecret = new Map()

        newValue.forEach((v, i, a) => {
          newSecret.set(v.secretKey, v.secretValue)
        })

        this.form.secrets = newSecret

        console.log(this.form.secrets)
      },
      deep: true
    }
  },
  computed: {
    newConfig () {
      if (this.config) {
        return !this.config.callbackid // should be false
      } else {
        return true
      }
    }
  },
  methods: {
    removeSecret (item) {
      var index = this.secrets.indexOf(item)
      if (index !== -1) {
        this.secrets.splice(index, 1)
      }
    },
    addSecret () {
      this.secrets.push({
        key: Date.now(),
        secretKey: '',
        secretValue: ''
      })
    },
    onSubmit () {
      this.$refs['form'].validate(valid => {
        if (valid) this.update()
      })
    },
    update () {
      this.client.setConfig(this.form).then(
        res => {
          this.$message({
            message: 'Config have been successfully updated',
            type: 'success'
          })
          this.showDialog = false
          this.$emit('updateConfig')
          if (this.newConfig) {
            this.$refs['form'].resetFields()
          }
        },
        err => {
          this.$message.error({
            message: 'Oops, error: ' + err
          })
        }
      )
    },
    title () {
      if (this.newConfig) {
        return 'Add config'
      } else {
        return 'Edit ' + this.form.title
      }
    },
    deleteConfig () {
      this.client.deleteConfig({ callbackid: this.form.callbackid }).then(
        res => {
          this.$message({
            message: 'Config have been successfully deleted',
            type: 'success'
          })
          this.showDialog = false
          this.$emit('updateConfig')
        },
        err => {
          this.$message.error({
            message: 'Oops, error: ' + err
          })
        }
      )
    }
  }
}
</script>

<style scoped>

</style>
