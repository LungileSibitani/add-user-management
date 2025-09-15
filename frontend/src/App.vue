<template>
  <div id="app" class="app-container">
    <h1 class="title">User Management System</h1>

    <!-- Add User Form -->
    <div class="card">
      <h2>Add New User</h2>
      <form @submit.prevent="addUser" class="form-container">
        <input v-model="newUser.email" type="email" placeholder="Enter Email" required />
        <input v-model="newUser.username" type="text" placeholder="Enter Username" required />
        <input v-model="newUser.first_name" type="text" placeholder="First Name" required />
        <input v-model="newUser.last_name" type="text" placeholder="Last Name" required />
        <input v-model="newUser.password_hash" type="password" placeholder="Password" required />
        <button type="submit" class="btn primary">➕ Add User</button>
      </form>
    </div>

    <!-- Toggle User List -->
    <button @click="showUsers = !showUsers" class="btn toggle-btn">
      {{ showUsers ? "⬆ Hide User List" : "⬇ Show User List" }}
    </button>

    <!-- User List -->
    <div v-if="showUsers" class="card user-list">
      <h2>User List</h2>
      <input
        v-model="search"
        type="text"
        placeholder="Search users by ID, username, email..."
        class="table-search"
        @input="resetAndFetch"
      />
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Email</th>
            <th>Username</th>
            <th>First</th>
            <th>Last</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in users" :key="user.id">
            <td>{{ user.id }}</td>
            <td>{{ user.email }}</td>
            <td>{{ user.username }}</td>
            <td>{{ user.first_name }}</td>
            <td>{{ user.last_name }}</td>
            <td>
              <button class="btn edit" @click="openEdit(user)">✏ Edit</button>
              <button class="btn delete" @click="deleteUser(user.id)">🗑 Delete</button>
            </td>
          </tr>
        </tbody>
      </table>

      <!-- Load More -->
      <button v-if="!loading && hasMore" @click="loadMore" class="btn toggle-btn">
        ⬇ Load More
      </button>
      <p v-if="loading">Loading...</p>
      <p v-if="!hasMore && users.length > 0">✅ All users loaded</p>
    </div>

    <!-- Edit Modal -->
    <div v-if="isEditing" class="modal-overlay">
      <div class="modal-card">
        <h2>Edit User</h2>
        <form @submit.prevent="updateUser" class="form-container">
          <input v-model="editUser.email" type="email" placeholder="Email" required />
          <input v-model="editUser.username" type="text" placeholder="Username" required />
          <input v-model="editUser.first_name" type="text" placeholder="First Name" required />
          <input v-model="editUser.last_name" type="text" placeholder="Last Name" required />
          <div class="modal-actions">
            <button type="submit" class="btn primary">✅ Update</button>
            <button type="button" class="btn cancel" @click="isEditing = false">✖ Cancel</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "App",
  data() {
    return {
      users: [],
      showUsers: false,
      search: "",
      newUser: {
        email: "",
        username: "",
        first_name: "",
        last_name: "",
        password_hash: "",
      },
      isEditing: false,
      editUser: {},

      // Pagination
      page: 1,
      limit: 50,
      hasMore: true,
      loading: false,
    };
  },
  methods: {
    async fetchUsers(reset = false) {
      try {
        this.loading = true;
        const res = await axios.get(
          `http://localhost:8080/api/users?page=${this.page}&limit=${this.limit}&q=${this.search}`
        );
        const data = res.data.data || res.data; // handle both formats

        if (reset) {
          this.users = data;
        } else {
          this.users = [...this.users, ...data];
        }

        if (data.length < this.limit) {
          this.hasMore = false; // no more data
        }
      } catch (err) {
        console.error("Failed to fetch users:", err);
      } finally {
        this.loading = false;
      }
    },
    resetAndFetch() {
      this.page = 1;
      this.hasMore = true;
      this.users = [];
      this.fetchUsers(true);
    },
    loadMore() {
      if (this.hasMore) {
        this.page++;
        this.fetchUsers();
      }
    },
    async addUser() {
      try {
        await axios.post("http://localhost:8080/api/users", this.newUser);
        this.newUser = { email: "", username: "", first_name: "", last_name: "", password_hash: "" };
        this.resetAndFetch();
      } catch (err) {
        console.error("Failed to add user:", err);
      }
    },
    async deleteUser(id) {
      try {
        await axios.delete(`http://localhost:8080/api/users/${id}`);
        this.resetAndFetch();
      } catch (err) {
        console.error("Failed to delete user:", err);
      }
    },
    openEdit(user) {
      this.editUser = { ...user };
      this.isEditing = true;
    },
    async updateUser() {
      try {
        await axios.put(`http://localhost:8080/api/users/${this.editUser.id}`, this.editUser);
        this.isEditing = false;
        this.resetAndFetch();
      } catch (err) {
        console.error("Failed to update user:", err);
      }
    },
  },
  mounted() {
    this.fetchUsers(true);
  },
};
</script>




<style>

.app-container {
  font-family: "Segoe UI", sans-serif;
  padding: 2rem;
  background: linear-gradient(135deg, #667eea, #764ba2);
  min-height: 100vh;
  color: #fff;
}

/* Title */
.title {
  text-align: center;
  margin-bottom: 2rem;
  font-size: 2rem;
  font-weight: bold;
}

/* Card style */
.card {
  background: rgba(255, 255, 255, 0.1);
  border-radius: 16px;
  padding: 1.5rem;
  margin-bottom: 1.5rem;
  backdrop-filter: blur(10px);
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.2);
}

/* Form */
.form-container {
  display: grid;
  gap: 0.8rem;
  max-width: 350px;
  margin: auto;
}

/* Input Fields - dark theme */
input {
  padding: 0.7rem;
  border: 1px solid rgba(255, 255, 255, 0.3);
  border-radius: 12px;
  outline: none;
  font-size: 0.95rem;
  background: rgba(255, 255, 255, 0.05);
  color: #fff;
  backdrop-filter: blur(5px);
  transition: all 0.3s ease;
}

input::placeholder {
  color: rgba(255, 255, 255, 0.6);
}

input:focus {
  border-color: #fff;
  box-shadow: 0 0 8px rgba(255, 255, 255, 0.6);
}

/* Buttons */
.btn {
  border: none;
  padding: 0.7rem 1.2rem;
  border-radius: 8px;
  cursor: pointer;
  font-size: 0.9rem;
  transition: all 0.3s ease;
}

/* Add User */
.btn.primary {
  background: linear-gradient(45deg, #43cea2, #185a9d);
  color: white;
}
.btn.primary:hover {
  opacity: 0.9;
}

/* Toggle User List */
.btn.toggle-btn {
  display: block;
  margin: 1rem auto;
  background: #ff9a9e;
  color: #333;
}

/* Edit button */
.btn.edit {
  background: linear-gradient(45deg, #6dd5ed, #2193b0);
  color: white;
  margin-right: 0.5rem;
}
.btn.edit:hover {
  opacity: 0.9;
}

/* Delete button */
.btn.delete {
  background: linear-gradient(45deg, #ff6a6a, #ff4e50);
  color: white;
}
.btn.delete:hover {
  opacity: 0.9;
}

/* Table */
.user-list table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 1rem;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 8px;
  overflow: hidden;
}
.user-list th,
.user-list td {
  padding: 0.8rem;
  text-align: left;
  border-bottom: 1px solid rgba(255, 255, 255, 0.2);
  color: #fff;
}
.user-list th {
  background: rgba(0, 0, 0, 0.2);
}

/* Table Search */
.table-search {
  width: 100%;
  padding: 0.7rem;
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.3);
  margin-bottom: 0.8rem;
  background: rgba(255, 255, 255, 0.05);
  color: #fff;
  outline: none;
  backdrop-filter: blur(5px);
  transition: all 0.3s ease;
}
.table-search::placeholder {
  color: rgba(255, 255, 255, 0.6);
}
.table-search:focus {
  border-color: #fff;
  box-shadow: 0 0 8px rgba(255, 255, 255, 0.6);
}

/* Modal */
.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}
.modal-card {
  background: rgba(255, 255, 255, 0.1);
  padding: 2rem;
  border-radius: 16px;
  backdrop-filter: blur(10px);
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.3);
  width: 90%;
  max-width: 400px;
}
.modal-actions {
  display: flex;
  gap: 0.5rem;
  margin-top: 1rem;
}
.btn.cancel {
  background: linear-gradient(45deg, #b0b0b0, #808080);
  color: white;
}
</style>
