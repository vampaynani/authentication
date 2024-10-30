<script setup lang="ts">
import { onMounted, reactive } from 'vue'
import { useUsers } from '@/services/users'

const { loadUsers, deleteUser, createUser, users } = useUsers()

const newUser = reactive({
  username: '',
  firstname: '',
  lastname: '',
  email: ''
})

onMounted(loadUsers)
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
    <form @submit.prevent="() => createUser(newUser)">
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
