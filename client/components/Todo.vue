<template>
  <div class="task-form">
    <v-form>
      <v-col sm="12" md="60">
        <v-text-field
          label="Task"
          name="task"
        />
      </v-col>
      <div class="add-btn">
        <v-btn @click="addTask()">
          追加
        </v-btn>
      </div>
    </v-form>
  </div>
</template>

<style>
.task-form {
    border: solid 3px #000000;
}
</style>

<script>
export default {
  data () {
    return {
      task: ''
    }
  },
  methods: {
    async addTask () {
      const formData = new FormData()
      formData.append('task', this.task)
      const response = await this.$axios.post('/api/addtask', formData, {
      }).catch((error) => {
        return error.response
      })
      console.info(response.data.message)
    }
  }
}
</script>
