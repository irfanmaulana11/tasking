import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Sidebar from './components/Sidebar';
import ProtectedRoute from './components/ProtectedRoute';
import Dashboard from './pages/Dashboard';
import TaskList from './pages/TaskList';
import CreateTask from './pages/CreateTask';
import EditTask from './pages/EditTask';
import UserManagement from './pages/UserManagement';
import Login from './pages/Login';
import Register from './pages/Register';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/*" element={
          <ProtectedRoute>
            <div className="flex min-h-screen bg-gray-100">
              <Sidebar />
              <div className="flex-1 ml-64">
                <Routes>
                  <Route path="/" element={<Dashboard />} />
                  <Route path="/tasks" element={<TaskList />} />
                  <Route path="/tasks/create" element={<CreateTask />} />
                  <Route path="/tasks/edit/:id" element={<EditTask />} />
                  <Route path="/users" element={<UserManagement />} />
                </Routes>
              </div>
            </div>
          </ProtectedRoute>
        } />
      </Routes>
    </Router>
  );
}

export default App;
