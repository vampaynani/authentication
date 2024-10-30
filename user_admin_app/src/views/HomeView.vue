<script setup lang="ts">
import { reactive } from 'vue'
import { useAuthStore } from '@/stores/auth'
const authStore = useAuthStore()

const loginInputs = reactive({
  username: '',
  password: ''
})

async function signIn() {
  const response = await fetch('http://localhost:3333/auth/login', {
    method: 'POST',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(loginInputs)
  })
  const data = await response.json()
  authStore.setSession(data.session, 'about')
}
</script>

<template>
  <main>
    <form @submit.prevent="signIn">
      <input v-model="loginInputs.username" type="text" placeholder="Username" />
      <input v-model="loginInputs.password" type="password" placeholder="Password" />
      <button>Iniciar sesión</button>
    </form>
  </main>
</template>
