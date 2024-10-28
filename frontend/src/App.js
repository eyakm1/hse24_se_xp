import React from 'react';
import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
import { useSelector } from 'react-redux';
import Login from './pages/Login';
import Dashboard from './pages/Dashboard';

function PrivateRoute({ children }) {
  const { isAuthenticated } = useSelector((state) => state.auth);
  return isAuthenticated ? children : <Navigate to="/login" />;
}

function App() {
  return (
  <Routes>
    <Route path="/login" element={<Login />} />
    <Route path="/dashboard" element={
      <PrivateRoute>
        <Dashboard />
      </PrivateRoute>
    } />
    <Route path="*" element={<Navigate to="/login" />} />
  </Routes>
  );
}

export default App;
