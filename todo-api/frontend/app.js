const API_URL = "http://localhost:8080/tasks";

// Fungsi untuk mengambil dan menampilkan tasks
async function fetchTasks() {
    try {
        const response = await fetch(API_URL);
        if (!response.ok) {
            throw new Error("Failed to fetch tasks");
        }
        const tasks = await response.json();
        const taskList = document.getElementById("taskList");
        taskList.innerHTML = "";

        tasks.forEach(task => {
            const taskCard = document.createElement("div");
            taskCard.className = "task-card";
            taskCard.innerHTML = `
                <div class="task-content">
                    <h5>${task.title}</h5>
                    <p>${task.description}</p>
                    <small>${task.completed ? "Completed" : "Not Completed"}</small>
                </div>
                <div class="task-actions">
                    <button class="btn btn-warning edit-btn" data-task='${JSON.stringify(task)}'>
                        <i class="fas fa-edit"></i> Edit
                    </button>
                    <button class="btn btn-danger delete-btn" data-task-id="${task.id}">
                        <i class="fas fa-trash"></i> Delete
                    </button>
                </div>
            `;
            taskList.appendChild(taskCard);
        });

        // Tambahkan event listener ke tombol edit dan delete
        document.querySelectorAll(".edit-btn").forEach(button => {
            button.addEventListener("click", function() {
                const task = JSON.parse(this.getAttribute("data-task"));
                openEditModal(task);
            });
        });

        document.querySelectorAll(".delete-btn").forEach(button => {
            button.addEventListener("click", function() {
                const taskId = this.getAttribute("data-task-id");
                deleteTask(taskId);
            });
        });
    } catch (error) {
        console.error("Error fetching tasks:", error);
    }
}

// Fungsi untuk menambahkan task
async function addTask(event) {
    event.preventDefault(); // Mencegah form submit default
    const title = document.getElementById("title").value;
    const description = document.getElementById("description").value;

    if (!title || !description) {
        alert("Title and description are required!");
        return;
    }

    try {
        const response = await fetch(API_URL, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ title, description }),
        });

        if (!response.ok) {
            throw new Error("Failed to add task");
        }

        fetchTasks(); // Refresh daftar task
        document.getElementById("taskForm").reset(); // Reset form
    } catch (error) {
        console.error("Error adding task:", error);
    }
}

// Fungsi untuk membuka modal edit
function openEditModal(task) {
    document.getElementById("editId").value = task.id;
    document.getElementById("editTitle").value = task.title;
    document.getElementById("editDescription").value = task.description;
    document.getElementById("editCompleted").checked = task.completed;

    // Tampilkan modal
    let editModal = new bootstrap.Modal(document.getElementById('editModal'));
    editModal.show();
}

// Fungsi untuk menyimpan perubahan setelah edit
async function saveEdit() {
    const id = document.getElementById("editId").value;
    const title = document.getElementById("editTitle").value;
    const description = document.getElementById("editDescription").value;
    const completed = document.getElementById("editCompleted").checked;

    try {
        const response = await fetch(`${API_URL}/${id}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ title, description, completed }),
        });

        if (!response.ok) {
            throw new Error("Failed to update task");
        }

        fetchTasks(); // Refresh daftar task
        let editModal = bootstrap.Modal.getInstance(document.getElementById('editModal'));
        editModal.hide(); // Tutup modal
    } catch (error) {
        console.error("Error updating task:", error);
    }
}

// Fungsi untuk menghapus task
async function deleteTask(id) {
    try {
        const response = await fetch(`${API_URL}/${id}`, {
            method: "DELETE",
        });

        if (!response.ok) {
            throw new Error("Failed to delete task");
        }

        fetchTasks(); // Refresh daftar task
    } catch (error) {
        console.error("Error deleting task:", error);
    }
}

// Event listener untuk form dan tombol save edit
document.getElementById("taskForm").addEventListener("submit", addTask);
document.getElementById("saveEditBtn").addEventListener("click", saveEdit);

// Ambil dan tampilkan tasks saat halaman dimuat
fetchTasks();