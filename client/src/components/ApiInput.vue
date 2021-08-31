<template>
  <div class="api-input-class">

    <el-row>
      <el-col span="1.5">
        <div> <span style="padding-left: 20px; font-size: 15px; font-weight: 600; color: gray">预设模板</span></div>
        <div><el-radio v-model="template" label="302" class="template-class">302跳转模板</el-radio></div>
        <div><el-radio v-model="template" label="ssrf" class="template-class">mock模板(JSON)</el-radio></div>
      </el-col>
      <el-col span="8">
        <el-form ref="form" :model="form" label-width="150px">
          <el-form-item label="Method: ">
            <el-radio-group v-model="apiPostForm.method" size="small">
              <el-radio-button label="GET"></el-radio-button>
              <el-radio-button label="POST"></el-radio-button>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="Path: ">
            <el-input v-model="apiPostForm.path" placeholder="请输入请求路径" class="status-code-class"></el-input>
          </el-form-item>
          <el-form-item label="StatusCode: " required>
            <el-input v-model="apiPostForm.status_code" placeholder="输入StatusCode" class="status-code-class"></el-input>
            <el-radio-group v-model="apiPostForm.status_code" class="status-code-radio">
              <el-radio :label="200">200</el-radio>
              <el-radio :label="302">302</el-radio>
              <el-radio :label="404">404</el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="Content-Type: " required>
            <el-input v-model="apiPostForm.content_type" placeholder="输入Content-Type" class="status-code-class"></el-input>
            <el-radio-group v-model="apiPostForm.content_type" class="status-code-radio">
              <el-radio label="text/html">text/html</el-radio>
              <el-radio label="application/json">application/json</el-radio>
              <el-radio label="text/plain">text/plain</el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="Header: ">
            <div>
              <el-input
                  placeholder="请输入key"
                  v-model="headerKey"
                  class="status-code-class">
              </el-input>
              :
              <el-input
                  placeholder="请输入value"
                  v-model="headerValue"
                  style="margin-left: 20px; width: 300px">
              </el-input>
              <el-button type="primary" icon="el-icon-plus" style="margin-left: 10px"></el-button>
            </div>

            <div>
              <el-input
                  placeholder="请输入复制的返回头"
                  type="textarea"
                  :rows="5"
                  style="width: 532px; margin-top: 15px"
                  v-model="headerText">
              </el-input>
            </div>
          </el-form-item>
          <el-form-item label="Body: " required>
            <div>
              <el-input
                  placeholder="请输入要返回的body"
                  type="textarea"
                  :rows="9"
                  style="width: 532px"
                  v-model="apiPostForm.body">
              </el-input>
            </div>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmit">立即创建</el-button>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "ApiInput",
  data() {
    return {
      template: '',
      headerKey: '',
      headerValue: '',
      headerText: '',
      apiPostForm: {
        method: 'GET',
        path: '',
        status_code: 200,
        content_type: 'text/html',
        body: ''
      }
    }
  },
  methods: {
    onSubmit() {
      console.log('submit!');
    }
  }
}
</script>

<style scoped>
.status-code-class {
  width: 200px;
}
.status-code-radio {
  margin-left: 30px;
}
.template-class {
  margin-top: 20px;
}
</style>