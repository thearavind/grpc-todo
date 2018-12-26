<template>
  <div id="app">
    <section>
      <span class="title-text">TODO gRPC Client</span>
      <div class="row justify-content-center mt-4">
        <input v-model="inputField" v-on:keyup.enter="addTodo" class="mr-1" placeholder="Todo Item">
        <button @click="addTodo" class="btn btn-primary">Add Todo</button>
      </div>
    </section>
    <section>
      <div class="row">
        <div class="offset-md-3 col-md-6 mt-3">
          <ul class="list-group justify-content-center">
            <li
              class="row list-group-item border mt-2 col-xs-1"
              v-for="todo in todos"
              v-bind:key="todo.id"
            >
              <div>
                <span>{{todo.task}}</span>
                <span @click="deleteTodo(todo)" class="offset-sm-1 col-sm-2 delete text-right">X</span>
              </div>
            </li>
          </ul>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import { addTodoParams, getTodoParams, deleteTodoParams } from "todo_pb";
import { todoServiceClient } from "todo_grpc_web_pb";
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";

export default {
  name: "app",
  components: {},
  data: function() {
    return {
      inputField: "",
      todos: []
    };
  },
  created: function() {
    this.client = new todoServiceClient("http://localhost:8080", null, null);
    this.getTodos();
  },
  methods: {
    getTodos: function() {
      let getRequest = new getTodoParams();
      this.client.getTodos(getRequest, {}, (err, response) => {
        this.todos = response.toObject().todosList;
        console.log(this.todos);
      });
    },
    addTodo: function() {
      let request = new addTodoParams();
      request.setTask(this.inputField);
      this.client.addTodo(request, {}, () => {
        this.inputField = "";
        this.getTodos();
      });
    },
    deleteTodo: function(todo) {
      let deleteRequest = new deleteTodoParams();
      deleteRequest.setId(todo.id);
      this.client.deleteTodo(deleteRequest, {}, (err, response) => {
        if (response.getMessage() === "success") {
          this.getTodos();
        }
      });
      console.log("todo -> ", todo.id);
    }
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.title-text {
  font-size: 22px;
}
</style>
