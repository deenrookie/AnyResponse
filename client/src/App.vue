<template>
  <div id="app">
    <el-row>
      <el-col :span="2">
        <div><span style="padding-left: 20px; font-size: 15px; font-weight: 600; color: gray">预设模板</span></div>
        <div>
          <el-radio v-model="template" label="302" class="template-class" @change="templateChange">302跳转模板</el-radio>
        </div>
        <div>
          <el-radio v-model="template" label="mock" class="template-class" @change="templateChange">mock模板(JSON)</el-radio>
        </div>
      </el-col>
      <el-col :span="10">
        <el-form ref="apiPostForm" :model="apiPostForm" label-width="150px">
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
            <el-input v-model.number="apiPostForm.status_code" placeholder="输入StatusCode"
                      class="status-code-class"></el-input>
            <el-radio-group v-model="apiPostForm.status_code" class="status-code-radio">
              <el-radio :label="200">200</el-radio>
              <el-radio :label="302">302</el-radio>
              <el-radio :label="404">404</el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="Content-Type: " required>
            <el-input v-model="apiPostForm.content_type" placeholder="输入Content-Type"
                      class="status-code-class"></el-input>
            <el-radio-group v-model="apiPostForm.content_type" class="status-code-radio">
              <el-radio label="text/html">text/html</el-radio>
              <el-radio label="application/json">application/json</el-radio>
              <el-radio label="text/plain">text/plain</el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="Header: ">
            <div>
              <el-autocomplete
                  placeholder="请输入key"
                  v-model="headerKey"
                  :fetch-suggestions="queryHeader"
                  :trigger-on-focus="true"
                  class="status-code-class">
              </el-autocomplete>
              :
              <el-input
                  placeholder="请输入value"
                  v-model="headerValue"
                  style="margin-left: 20px; width: 20vh">
              </el-input>
              <el-button type="primary" icon="el-icon-plus" style="margin-left: 10px" @click="addHeader"></el-button>
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
            <el-button type="primary" @click="addApi">立即创建</el-button>
          </el-form-item>
        </el-form>
      </el-col>
      <el-col :span="6" style="margin-left: 40px">
        <div v-for="(item, pos) in apiList" :key="pos">
          <span style="width: 45px;display: inline-block"><el-tag type="success" effect="dark"
                                                                  size="medium"> {{ item.method }} </el-tag> </span>
          <el-tag style="width: 330px;margin-left: 20px; font-size: 16px" effect="dark"
                  size="medium" class="api-tag-class">
            {{ item.path }}
          </el-tag>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>

export default {
  name: 'App',
  data() {
    return {
      template: '',
      headerKey: '',
      headerValue: '',
      headerText: '',
      apiPostForm: {
        method: 'GET',
        path: '',
        headers: {},
        status_code: 200,
        content_type: 'text/html',
        body: ''
      },
      apiList: [
        {
          method: 'GET',
          path: "/api/example/ping",
        },
        {
          method: 'POST',
          path: "/api/example/ping2",
        }
      ],
      allHeaders: [
        { "value": "Set-Cookie" },
        { "value": "Content-Type" },
        {"value": "Access-Control-Allow-Origin"},
        {"value": "Content-Encoding"},
        {"value": "Content-Length"},
        {"value": "Location"}
      ]
    }
  },
  created() {
    this.getApiList()
  },
  methods: {
    createFilter(queryString) {
      return (restaurant) => {
        return (restaurant.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0);
      };
    },
    queryHeader(queryString, cb) {
      const headers = this.allHeaders;
      const results = queryString ? headers.filter(this.createFilter(queryString)) : headers;
      cb(results);
    },
    addHeader() {
      if(this.headerKey && this.headerValue) {
        this.apiPostForm.headers[this.headerKey] = this.headerValue
        this.headerText = JSON.stringify(this.apiPostForm.headers)
      }
    },
    getApiList() {
      this.axios.get('http://localhost:9999/api/list').then((response) => {
        console.log(response)
        this.apiList = response.data;
      }).catch((response) => {
        console.log(response);
      })
    },
    addApi() {
      this.axios.post('http://localhost:9999/api/add', this.apiPostForm)
          .then(res => {
            console.log('res=>', res)
            this.$message.success("提交成功")
            this.getApiList()
          })
    },
    templateChange(label) {
      if (label === '302') {
        this.apiPostForm.status_code  = 302
        this.apiPostForm.headers["Location"] = "https://www.example.com"
        this.headerText = JSON.stringify(this.apiPostForm.headers)
      }
      if (label === 'mock') {
        this.apiPostForm.status_code  = 200
        this.apiPostForm.content_type = 'application/json'
        this.headerText = JSON.stringify(this.apiPostForm.headers)
      }
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  margin-top: 15vh;
  margin-left: 15vh;
}

.status-code-class {
  width: 20vh;
}

.status-code-radio {
  margin-left: 30px;
}

.template-class {
  margin-top: 20px;
}

.api-tag-class {
  margin: 10px;
  font-size: 15px;
  font-weight: 400;
}
</style>
