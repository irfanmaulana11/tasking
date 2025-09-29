import React, { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';

interface TaskHistory {
  action_by: string;
  action: string;
  note: string;
  created_at: string;
}

interface Task {
  id: string;
  title: string;
  description: string;
  Assignee: string;
  assigned_leader: string;
  status: string;
  progress: number;
  progress_by: string;
  created_by: string;
  created_at: string;
  updated_at: string;
  task_history: TaskHistory[];
}

const TaskList: React.FC = () => {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [loading, setLoading] = useState(true);
  const [selectedTask, setSelectedTask] = useState<Task | null>(null);
  const [showHistory, setShowHistory] = useState(false);
  const [showProgressModal, setShowProgressModal] = useState(false);
  const [progressValue, setProgressValue] = useState(0);
  const [actionLoading, setActionLoading] = useState<string | null>(null);
  const [viewMode, setViewMode] = useState<'card' | 'board'>('card');
  const navigate = useNavigate();
  
  const userData = JSON.parse(localStorage.getItem('user') || '{}');
  const userRole = userData.role || '';

  const statusColors = {
    'Submitted': 'bg-blue-100 text-blue-800 border-blue-200',
    'Revision': 'bg-yellow-100 text-yellow-800 border-yellow-200',
    'Approved': 'bg-green-100 text-green-800 border-green-200',
    'In Progress': 'bg-purple-100 text-purple-800 border-purple-200',
    'Completed': 'bg-gray-100 text-gray-800 border-gray-200'
  };

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

  const groupedTasks = tasks.reduce((groups: { [key: string]: Task[] }, task) => {
    const status = task.status;
    if (!groups[status]) {
      groups[status] = [];
    }
    groups[status].push(task);
    return groups;
  }, {});

  const showTaskHistory = (task: Task) => {
    setSelectedTask(task);
    setShowHistory(true);
  };

  const handleTaskAction = async (taskId: string, action: string, progressValue?: number) => {
    setActionLoading(taskId + action);
    try {
      const token = localStorage.getItem('token');
      let url = '';
      let body = {};
      
      switch (action) {
        case 'progress':
          url = `http://localhost:8080/api/tasks/${taskId}/progress`;
          body = { progress: progressValue };
          break;
        case 'override':
          url = `http://localhost:8080/api/tasks/${taskId}/progress/overide`;
          body = { progress: progressValue };
          break;
        case 'revise':
          url = `http://localhost:8080/api/tasks/${taskId}/revise`;
          break;
        case 'approve':
          url = `http://localhost:8080/api/tasks/${taskId}/approve`;
          break;
        case 'complete':
          url = `http://localhost:8080/api/tasks/${taskId}/complete`;
          break;
      }
      
      const response = await fetch(url, {
        method: 'PATCH',
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(body)
      });
      
      if (response.ok) {
        fetchTasks(); // Refresh tasks
        setShowProgressModal(false);
      } else {
        const data = await response.json();
        alert(data.error || 'Action failed');
      }
    } catch (error) {
      alert('Connection error');
    } finally {
      setActionLoading(null);
    }
  };

  const openProgressModal = (task: Task, action: string) => {
    setSelectedTask(task);
    setProgressValue(task.progress);
    setShowProgressModal(true);
  };

  const getAvailableActions = (task: Task) => {
    const actions = [];
    const userData = JSON.parse(localStorage.getItem('user') || '{}');
    const currentUser = userData.username || userData.user_name || userData.name || '';
    
    // Progress update - for pelaksana on their own tasks
    if (userRole === 'pelaksana' && (task.status === 'Approved' || task.status === 'In Progress')) {
      actions.push({
        label: 'Update Progress',
        action: 'progress',
        color: 'bg-blue-500 hover:bg-blue-600',
        needsProgress: true
      });
    }
    
    // Leader actions
    if (userRole === 'leader') {
      if (task.status === 'Submitted') {
        actions.push(
          {
            label: 'Approve',
            action: 'approve',
            color: 'bg-green-500 hover:bg-green-600',
            needsProgress: false
          },
          {
            label: 'Revise',
            action: 'revise',
            color: 'bg-yellow-500 hover:bg-yellow-600',
            needsProgress: false
          }
        );
      }
      
      // Leader can approve task in revision status if assigned to them
      if (task.status === 'Revision' && task.Assignee === currentUser) {
        actions.push({
          label: 'Approve',
          action: 'approve',
          color: 'bg-green-500 hover:bg-green-600',
          needsProgress: false
        });
      }
      
      if (task.status === 'In Progress') {
        actions.push({
          label: 'Override Progress',
          action: 'override',
          color: 'bg-purple-500 hover:bg-purple-600',
          needsProgress: true
        });
      }
    }
    
    // Complete action - for both leader and pelaksana on In Progress tasks
    if ((userRole === 'leader' || userRole === 'pelaksana') && task.status === 'In Progress') {
      actions.push({
        label: 'Complete',
        action: 'complete',
        color: 'bg-gray-500 hover:bg-gray-600',
        needsProgress: false
      });
    }
    
    return actions;
  };

  if (loading) {
    return (
      <div className="p-6">
        <div className="text-center">Loading tasks...</div>
      </div>
    );
  }

  return (
    <div className="p-6">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-3xl font-bold text-gray-800">Task List</h1>
        <div className="flex items-center space-x-3">
          {/* View Mode Toggle */}
          <div className="flex bg-gray-200 rounded-lg p-1">
            <button
              onClick={() => setViewMode('card')}
              className={`px-3 py-1 rounded-md text-sm transition-colors ${
                viewMode === 'card' 
                  ? 'bg-white text-gray-900 shadow-sm' 
                  : 'text-gray-600 hover:text-gray-900'
              }`}
            >
              📋 Card View
            </button>
            <button
              onClick={() => setViewMode('board')}
              className={`px-3 py-1 rounded-md text-sm transition-colors ${
                viewMode === 'board' 
                  ? 'bg-white text-gray-900 shadow-sm' 
                  : 'text-gray-600 hover:text-gray-900'
              }`}
            >
              📏 Board View
            </button>
          </div>
          
          <button
            onClick={fetchTasks}
            className="px-4 py-2 bg-gray-600 text-white rounded-md hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500"
          >
            🔄 Refresh
          </button>
          <Link
            to="/tasks/create"
            className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            + Create Task
          </Link>
        </div>
      </div>
      
      {Object.keys(groupedTasks).length === 0 ? (
        <div className="bg-white rounded-lg shadow-md p-6">
          <p className="text-gray-500 text-center">No tasks found.</p>
        </div>
      ) : viewMode === 'card' ? (
        <div className="space-y-8">
          {Object.entries(groupedTasks).map(([status, statusTasks]) => (
            <div key={status}>
              <h2 className="text-xl font-semibold text-gray-700 mb-4">
                {status} ({statusTasks.length})
              </h2>
              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                {statusTasks.map((task) => (
                  <TaskCard 
                    key={task.id} 
                    task={task} 
                    onShowHistory={showTaskHistory}
                    onActionClick={(taskId, action, needsProgress) => {
                      if (needsProgress) {
                        openProgressModal(task, action);
                      } else {
                        handleTaskAction(taskId, action);
                      }
                    }}
                    onEditClick={(taskId) => navigate(`/tasks/edit/${taskId}`)}
                    getAvailableActions={getAvailableActions}
                    actionLoading={actionLoading}
                    userRole={userRole}
                    statusColors={statusColors}
                  />
                ))}
              </div>
            </div>
          ))}
        </div>
      ) : (
        <div className="flex gap-6 overflow-x-auto pb-6">
          {['Submitted', 'Revision', 'Approved', 'In Progress', 'Completed'].map((status) => {
            const statusTasks = groupedTasks[status] || [];
            return (
              <div key={status} className="flex-shrink-0 w-80">
                <div className="bg-gray-100 rounded-lg p-4">
                  <div className="flex items-center justify-between mb-4">
                    <h3 className="font-semibold text-gray-800">{status}</h3>
                    <span className={`px-2 py-1 rounded-full text-xs font-medium border ${statusColors[status as keyof typeof statusColors] || 'bg-gray-100 text-gray-800'}`}>
                      {statusTasks.length}
                    </span>
                  </div>
                  
                  <div className="space-y-3 max-h-96 overflow-y-auto">
                    {statusTasks.map((task) => (
                      <TaskCard 
                        key={task.id} 
                        task={task} 
                        compact 
                        onShowHistory={showTaskHistory}
                        onActionClick={(taskId, action, needsProgress) => {
                          if (needsProgress) {
                            openProgressModal(task, action);
                          } else {
                            handleTaskAction(taskId, action);
                          }
                        }}
                        onEditClick={(taskId) => navigate(`/tasks/edit/${taskId}`)}
                        getAvailableActions={getAvailableActions}
                        actionLoading={actionLoading}
                        userRole={userRole}
                        statusColors={statusColors}
                      />
                    ))}
                    {statusTasks.length === 0 && (
                      <div className="text-center text-gray-500 py-8">
                        No tasks
                      </div>
                    )}
                  </div>
                </div>
              </div>
            );
          })}
        </div>
      )}

      {/* History Modal */}
      {showHistory && selectedTask && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 max-w-2xl w-full mx-4 max-h-96 overflow-y-auto">
            <div className="flex justify-between items-center mb-4">
              <h3 className="text-lg font-semibold">Task History: {selectedTask.title}</h3>
              <button
                onClick={() => setShowHistory(false)}
                className="text-gray-500 hover:text-gray-700"
              >
                ✕
              </button>
            </div>
            
            {selectedTask.task_history.length === 0 ? (
              <p className="text-gray-500">No history available.</p>
            ) : (
              <div className="space-y-3">
                {selectedTask.task_history.map((history, index) => (
                  <div key={index} className="border-l-4 border-blue-500 pl-4 py-2">
                    <div className="flex justify-between items-start">
                      <div>
                        <div className="font-medium text-gray-800">{history.action}</div>
                        <div className="text-sm text-gray-600">by {history.action_by}</div>
                        {history.note && (
                          <div className="text-sm text-gray-500 mt-1">{history.note}</div>
                        )}
                      </div>
                      <div className="text-xs text-gray-400">
                        {new Date(history.created_at).toLocaleString()}
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </div>
        </div>
      )}

      {/* Progress Modal */}
      {showProgressModal && selectedTask && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 max-w-md w-full mx-4">
            <div className="flex justify-between items-center mb-4">
              <h3 className="text-lg font-semibold">Update Progress: {selectedTask.title}</h3>
              <button
                onClick={() => setShowProgressModal(false)}
                className="text-gray-500 hover:text-gray-700"
              >
                ✕
              </button>
            </div>
            
            <div className="mb-4">
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Progress: {progressValue}%
              </label>
              <input
                type="range"
                min="0"
                max="100"
                value={progressValue}
                onChange={(e) => setProgressValue(parseInt(e.target.value))}
                className="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer"
              />
              <div className="flex justify-between text-xs text-gray-500 mt-1">
                <span>0%</span>
                <span>50%</span>
                <span>100%</span>
              </div>
            </div>
            
            <div className="flex justify-end space-x-3">
              <button
                onClick={() => setShowProgressModal(false)}
                className="px-4 py-2 border border-gray-300 text-gray-700 rounded hover:bg-gray-50"
              >
                Cancel
              </button>
              <button
                onClick={() => {
                  const action = userRole === 'leader' ? 'override' : 'progress';
                  handleTaskAction(selectedTask.id, action, progressValue);
                }}
                disabled={actionLoading !== null}
                className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 disabled:opacity-50"
              >
                {actionLoading ? 'Updating...' : 'Update Progress'}
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

// TaskCard Component
interface TaskCardProps {
  task: Task;
  compact?: boolean;
  onShowHistory: (task: Task) => void;
  onActionClick: (taskId: string, action: string, needsProgress?: boolean) => void;
  onEditClick: (taskId: string) => void;
  getAvailableActions: (task: Task) => any[];
  actionLoading: string | null;
  userRole: string;
  statusColors: any;
}

const TaskCard: React.FC<TaskCardProps> = ({ 
  task, 
  compact = false, 
  onShowHistory, 
  onActionClick, 
  onEditClick, 
  getAvailableActions, 
  actionLoading, 
  userRole, 
  statusColors 
}) => {
  return (
    <div className={`bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow ${
      task.status === 'Revision' ? 'border-l-4 border-yellow-500' : ''
    } ${compact ? 'p-4' : 'p-6'}`}>
      <div className="flex justify-between items-start mb-3">
        <h3 className={`font-semibold text-gray-800 truncate ${compact ? 'text-sm' : 'text-lg'}`}>
          {task.title}
        </h3>
        {!compact && (
          <span className={`px-2 py-1 rounded-full text-xs font-medium border ${statusColors[task.status as keyof typeof statusColors] || 'bg-gray-100 text-gray-800'}`}>
            {task.status}
            {task.status === 'Revision' && userRole === 'pelaksana' && (
              <span className="ml-1 text-yellow-600">⚠️</span>
            )}
          </span>
        )}
      </div>
      
      <p className={`text-gray-600 mb-3 line-clamp-2 ${compact ? 'text-xs' : 'text-sm'}`}>
        {task.description}
      </p>
      
      <div className={`space-y-2 text-gray-500 ${compact ? 'text-xs' : 'text-sm'}`}>
        <div>Assignee: {task.Assignee}</div>
        {!compact && <div>Leader: {task.assigned_leader}</div>}
        <div className="flex items-center">
          <span>Progress: {task.progress}%</span>
          <div className="ml-2 flex-1 bg-gray-200 rounded-full h-2">
            <div 
              className="bg-blue-600 h-2 rounded-full transition-all duration-300" 
              style={{ width: `${task.progress}%` }}
            ></div>
          </div>
        </div>
        {!compact && <div>Created: {new Date(task.created_at).toLocaleDateString()}</div>}
      </div>
      
      <div className="mt-4">
        {!compact && (
          <div className="flex flex-wrap gap-2 mb-3">
            {getAvailableActions(task).map((actionItem, actionIndex) => (
              <button
                key={actionIndex}
                onClick={() => onActionClick(task.id, actionItem.action, actionItem.needsProgress)}
                disabled={actionLoading === task.id + actionItem.action}
                className={`px-2 py-1 text-white text-xs rounded transition-colors disabled:opacity-50 ${actionItem.color}`}
              >
                {actionLoading === task.id + actionItem.action ? 'Loading...' : actionItem.label}
              </button>
            ))}
          </div>
        )}
        
        <div className="flex justify-between items-center">
          <div className="flex space-x-2">
            <button
              onClick={() => onShowHistory(task)}
              className={`bg-gray-500 text-white rounded hover:bg-gray-600 transition-colors ${
                compact ? 'px-2 py-1 text-xs' : 'px-3 py-1 text-sm'
              }`}
            >
              {compact ? 'History' : 'View History'}
            </button>
            {task.status === 'Revision' && userRole === 'pelaksana' && (
              <button
                onClick={() => onEditClick(task.id)}
                className={`bg-orange-500 text-white rounded hover:bg-orange-600 transition-colors ${
                  compact ? 'px-2 py-1 text-xs' : 'px-3 py-1 text-sm'
                }`}
              >
                Edit
              </button>
            )}
          </div>
          {!compact && (
            <div className="text-xs text-gray-400">
              Updated: {new Date(task.updated_at).toLocaleDateString()}
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default TaskList;