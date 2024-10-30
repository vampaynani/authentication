import type { User } from '@/types'
import { apiCall } from '@/services/utils'
import { ref } from 'vue'

export function useUsers() {
  const users = ref<User[]>([])

  async function loadUsers() {
    const data = await apiCall('/users')
    users.value = data as User[]
  }

  async function createUser(newUser: {
    username: string
    firstname: string
    lastname: string
    email: string
  }) {
    const createdUser = await apiCall('/users', {
      method: 'POST',
      data: newUser
    })
    users.value = [...users.value, createdUser]
  }

  async function deleteUser(id: number) {
    const deleted = await apiCall(`/users/${id}`, {
      method: 'DELETE'
    })
    if (deleted) {
      users.value = users.value.filter((user) => user.id !== id)
    }
  }

  return { loadUsers, createUser, deleteUser, users }
}
