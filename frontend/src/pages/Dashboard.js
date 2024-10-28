import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { fetchAssignments } from '../store/slices/assignmentsSlice';
import AssignmentList from '../components/AssignmentList';

function Dashboard() {
  const dispatch = useDispatch();
  const { items: assignments, loading, error } = useSelector((state) => state.assignments);

  useEffect(() => {
    dispatch(fetchAssignments());
  }, [dispatch]);

  return (
    <div>
      <h1>Dashboard</h1>
      {loading && <p>Loading assignments...</p>}
      {error && <p>Error loading assignments: {error}</p>}
      {!loading && !error && <AssignmentList assignments={assignments} />}
    </div>
  );
}

export default Dashboard;
