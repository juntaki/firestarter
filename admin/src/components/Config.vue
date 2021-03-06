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

      <el-form-item label="Text" prop="texttemplate"
      :rules="[{ required: true, message: 'Please input Text Template', trigger: 'change' }]">
        <el-input v-model="form.texttemplate" placeholder="Deploy my app"></el-input>
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
          <el-col :span="10"><el-input v-model="secret.secretKey" placeholder="AWS_TOKEN" :disabled="secret.disabled"></el-input></el-col>
          <el-col :span="10"><el-input v-model="secret.secretValue" placeholder="AKIAXXXXX" :disabled="secret.disabled"></el-input></el-col>
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
import twirp from '../../proto/config_pb_twirp'
import pb from '../../proto/config_pb'
export default {
  props: ['config', 'channels'],
  data () {
    const form = this.config
      ? JSON.parse(JSON.stringify(this.config))
      : {
        id: null
      }

    const secrets = []
    if (this.config) {
      // Set temporary secrets
      this.config.secretsList.forEach(secret => {
        secrets.push({
          key: secrets.length,
          disabled: true,
          secretKey: secret.key,
          secretValue: secret.value
        })
      })
    }

    const host = location.protocol + '//' + location.host
    return {
      client: twirp.createConfigServiceClient(host),
      showDialog: false,
      showDeleteDialog: false,
      form: form,
      secrets: secrets,
      urlTemplatePlaceholder:
        'https://example.com/deploy?param={{index .matched 1}}&value={{value}}',
      bodyTemplatePlaceholder: "{ value: '{{value}}' }"
    }
  },
  computed: {
    newConfig () {
      if (this.config) {
        return !this.config.id // should be false
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
        key: Date.now(), // just for key for vue
        disabled: false,
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
      const config = new pb.Config()
      config.setId(this.form.id)
      config.setTitle(this.form.title)
      config.setChannelsList(this.form.channelsList)
      config.setRegexp(this.form.regexp)
      config.setTexttemplate(this.form.texttemplate)
      config.setActionsList(this.form.actionsList)
      config.setConfirm(this.form.confirm)
      config.setUrltemplate(this.form.urltemplate)
      config.setBodytemplate(this.form.bodytemplate)
      config.setSecretsList([])
      this.secrets.forEach((v, i, a) => {
        const pbsec = new pb.Secret()
        pbsec.setKey(v.secretKey)
        pbsec.setValue(v.secretValue)
        config.addSecrets(pbsec)
      })

      this.client.setConfigRaw(config).then(
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
      this.client.deleteConfig({ id: this.form.id }).then(
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
