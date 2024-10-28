import React from 'react';
import { Route, Routes, Navigate } from 'react-router-dom';
import { useSelector } from 'react-redux';
import Login from './pages/Login';
import Dashboard from './pages/Dashboard';
import AssignmentDetail from './pages/AssignmentDetail';
import Submission from './pages/Submission';
import TeacherDashboard from './pages/TeacherDashboard';
import AssignmentSubmissions from './pages/AssignmentSubmissions';

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
      <Route path="/assignments/:id" element={
        <PrivateRoute>
          <AssignmentDetail />
        </PrivateRoute>
      } />
      <Route path="/assignments/:id/submit" element={
        <PrivateRoute>
          <Submission />
        </PrivateRoute>
      } />
       <Route path="/teacher-dashboard" element={
        <PrivateRoute>
            <TeacherDashboard />
        </PrivateRoute>
      } />
        <Route path="/teacher-dashboard/assignments/:id/submissions" element={
        <PrivateRoute>
            <AssignmentSubmissions />
        </PrivateRoute>
      } />
      <Route path="*" element={<Navigate to="/login" />} />
    </Routes>
  );
}

export default App;
