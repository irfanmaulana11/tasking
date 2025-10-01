import React, { useState, useEffect } from 'react';

interface Task {
  id: string;
  title: string;
  status: string;
}

const Dashboard: React.FC = () => {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [loading, setLoading] = useState(true);

  const fetchTasks = async () => {
    try {
      const token = localStorage.getItem('token');
      const response = await fetch('http://localhost:8080/api/tasks/', {
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json'
        }
      });
      
      if (response.ok) {
        const data = await response.json();
        setTasks(data.data || []);
      }
    } catch (error) {
      console.error('Error fetching tasks:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchTasks();
  }, []);

  const getTaskCountByStatus = (status: string) => {
    return tasks.filter(task => task.status === status).length;
  };

  const taskStats = [
    { title: 'Submitted', count: getTaskCountByStatus('Submitted'), color: 'bg-blue-500'},
    { title: 'Revision', count: getTaskCountByStatus('Revision'), color: 'bg-yellow-500' },
    { title: 'Approved', count: getTaskCountByStatus('Approved'), color: 'bg-green-500'},
    { title: 'In Progress', count: getTaskCountByStatus('In Progress'), color: 'bg-purple-500'},
    { title: 'Completed', count: getTaskCountByStatus('Completed'), color: 'bg-gray-500'}
  ];

  if (loading) {
    return (
      <div className="p-6">
        <div className="text-center">Loading dashboard...</div>
      </div>
    );
  }

  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold text-gray-800 mb-8">Dashboard</h1>
      
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {taskStats.map((stat, index) => (
          <div key={index} className="bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow">
            <div className="flex items-center mb-4">
              <div className={`w-12 h-12 ${stat.color} rounded-lg flex items-center justify-center`}>
                <span className="text-white text-xl font-bold">{stat.count}</span>
              </div>
              <div className="ml-4">
                <h3 className="text-lg font-semibold text-gray-800">{stat.title}</h3>
              </div>
            </div>
          </div>
        ))}
      </div>
      
      <div className="mt-8 bg-white rounded-lg shadow-md p-6">
        <h2 className="text-xl font-semibold text-gray-800 mb-4">Summary</h2>
        <div className="text-gray-600">
          <p>Total Tasks: <span className="font-semibold">{tasks.length}</span></p>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;