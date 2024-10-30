<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'

interface User {
  id: number
  username: string
  firstname: string
  lastname: string
  email: string
}

const newUser = reactive({
  username: '',
  firstname: '',
  lastname: '',
  email: ''
})

const users = ref<User[]>([])

async function loadUsers() {
  const response = await fetch('http://localhost:3333/users', {
    credentials: 'include'
  })
  const data = await response.json()
  users.value = data as User[]
}

async function createUser() {
  const response = await fetch('http://localhost:3333/users', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(newUser)
  })
  const data = await response.json()
  users.value = [...users.value, data]
}

async function deleteUser(id: number) {
  const response = await fetch(`http://localhost:3333/users/${id}`, {
    method: 'DELETE'
  })
  if (response.ok) {
    users.value = users.value.filter((user) => user.id !== id)
  }
}

onMounted(() => {
  loadUsers()
})
</script>
<template>
  <section>
    <h1>Users</h1>
    <ul>
      <li v-for="user in users" :key="user.id">
        {{ user.id }}. {{ user.username }} - {{ user.firstname }} {{ user.lastname }}
        <button @click="() => deleteUser(user.id)">Borrar</button>
      </li>
    </ul>
    <h2>Add User</h2>
    <form @submit.prevent="createUser">
      <input v-model="newUser.username" type="text" placeholder="Username" />
      <input v-model="newUser.firstname" type="text" placeholder="Firstname" />
      <input v-model="newUser.lastname" type="text" placeholder="Lastname" />
      <input v-model="newUser.email" type="email" placeholder="Email" />
      <button type="submit">Add User</button>
    </form>
  </section>
</template>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
