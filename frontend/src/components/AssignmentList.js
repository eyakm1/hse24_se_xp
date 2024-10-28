import React from 'react';
import { Link } from 'react-router-dom';

function AssignmentList({ assignments }) {
  return (
    <div>
      <h2>Assignments</h2>
      <ul>
        {assignments.map((assignment) => (
          <li key={assignment.id}>
            <h3>{assignment.title}</h3>
            <p>{assignment.description}</p>
            <p>Due: {assignment.dueDate}</p>
            <Link to={`/assignments/${assignment.id}`}>View Details</Link>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default AssignmentList;
