<template>
  <div>
    <el-card v-for="config in configList" :key="config.callbackid" class="config-card">
      <div slot="header" class="clearfix">
        <span style="font-size: 30px">{{config.title}}</span>
        <config @updateConfig="update()" :config="config" :channels="channels" style="float: right" />
      </div>
      <el-row>
        <el-col :span="6">ID(Auto-assigned)</el-col>
        <el-col :span="18">{{config.callbackid}}</el-col>
      </el-row>
      <el-row>
        <el-col :span="6">Channels</el-col>
        <el-col :span="18">{{config.channelsList.join(',')}}</el-col>
      </el-row>
      <el-row>
        <el-col :span="6">Regexp</el-col>
        <el-col :span="18">{{config.regexp}}</el-col>
      </el-row>
      <el-row>
        <el-col :span="6">Text</el-col>
        <el-col :span="18">{{config.text}}</el-col>
      </el-row>
      <el-row>
        <el-col :span="6">Actions</el-col>
        <el-col :span="18">{{config.actionsList.join(',')}}</el-col>
      </el-row>
      <el-row>
        <el-col :span="6">Confirm</el-col>
        <el-col :span="18">{{config.confirm}}</el-col>
      </el-row>
      <el-row>
        <el-col :span="6">URL Template</el-col>
        <el-col :span="18">{{config.urltemplate}}</el-col>
      </el-row>
      <el-row>
        <el-col :span="6">Body Template</el-col>
        <el-col :span="18">{{config.bodytemplate}}</el-col>
      </el-row>
    </el-card>
    <div class="config-card">
      <config @updateConfig="update()" :channels="channels"/>
    </div>
  </div>
</template>

<script>
import config from '../../proto/config_pb_twirp'
import pb from '../../proto/config_pb'
import Config from './Config'
import DeleteConfig from './DeleteConfig'
export default {
  components: {
    Config,
    DeleteConfig
  },
  data () {
    const host = location.protocol + '//' + location.host
    const client = config.createConfigServiceClient(host)
    return {
      configList: [],
      client: client,
      channels: []
    }
  },
  mounted () {
    this.updateChannel()
    this.update()
  },
  methods: {
    update () {
      const c = new pb.GetConfigListRequest()
      this.client.getConfigList(c).then(
        res => {
          this.configList = res.configList
        },
        () => {}
      )
    },
    updateChannel () {
      const r = new pb.GetChannelsRequest()
      this.client.getChannels(r).then(
        res => {
          this.channels = res.listList
        },
        () => {}
      )
    }
  }
}
</script>

<style scoped>
.el-row {
  margin-bottom: 10px;
}

.clearfix:after {
  clear: both;
}

.config-card {
  margin: 20px;
}
</style>
