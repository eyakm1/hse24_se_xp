import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { fetchAssignments } from '../store/slices/assignmentsSlice';
import { Link } from 'react-router-dom';

function TeacherDashboard() {
  const dispatch = useDispatch();
  const { items: assignments, loading, error } = useSelector((state) => state.assignments);

  useEffect(() => {
    dispatch(fetchAssignments());
  }, [dispatch]);

  return (
    <div>
      <h1>Teacher Dashboard</h1>
      {loading && <p>Loading assignments...</p>}
      {error && <p>Error loading assignments: {error}</p>}
      <ul>
        {assignments.map((assignment) => (
          <li key={assignment.id}>
            <h3>{assignment.title}</h3>
            <p>Due Date: {assignment.dueDate}</p>
            <Link to={`/teacher-dashboard/assignments/${assignment.id}/submissions`}>View Submissions</Link>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default TeacherDashboard;
