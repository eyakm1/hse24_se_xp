import React from 'react';
import { useDispatch } from 'react-redux';
import { logout } from '../store/slices/authSlice';

function Dashboard() {
  const dispatch = useDispatch();

  const handleLogout = () => {
    dispatch(logout());
    window.location.href = '/login';
  };

  return (
    <div>
      <h1>Welcome to the Dashboard</h1>
      <p>This is where assignments, submissions, and grades will be managed.</p>
      <button onClick={handleLogout}>Logout</button>
    </div>
  );
}

export default Dashboard;
