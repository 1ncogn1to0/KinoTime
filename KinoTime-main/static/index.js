const apiClient = {
    async getAllUsers() {
        const response = await fetch('/users');
        return response.json();
    },

    async getUserById(id) {
        const response = await fetch(`/users/${id}`);
        return response.json();
    },

    async createUser(user) {
        const response = await fetch('/users', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(user),
        });
        return response.json();
    },

    async updateUser(id, user) {
        const response = await fetch(`/users/${id}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(user),
        });
        return response.json();
    },

    async deleteUser(id) {
        const response = await fetch(`/users/${id}`, {
            method: 'DELETE',
        });
        return response.json();
    },
};
