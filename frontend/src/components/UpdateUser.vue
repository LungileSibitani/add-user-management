<template>
  <div>
    <h2>Update User</h2>
    <form @submit.prevent="updateUser">
      <input v-model="id" placeholder="User ID" />
      <input v-model="username" placeholder="Username" />
      <input v-model="email" placeholder="Email" />
      <button type="submit">Update</button>
    </form>
  </div>
</template>

<script lang="ts">
import { ref } from 'vue'

export default {
  setup() {
    const id = ref('')
    const username = ref('')
    const email = ref('')

    const updateUser = async () => {
      try {
        await fetch(`http://localhost:8080/users/${id.value}`, {
          method: 'PUT',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ username: username.value, email: email.value })
        })
        alert('User updated!')
      } catch (err) {
        console.error(err)
      }
    }

    return { id, username, email, updateUser }
  }
}
</script>

