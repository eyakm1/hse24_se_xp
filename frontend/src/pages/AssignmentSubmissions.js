import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { fetchSubmissionsForAssignment, gradeSubmission } from '../store/slices/assignmentsSlice';

function AssignmentSubmissions() {
  const { id } = useParams();
  const dispatch = useDispatch();
  const { submissions, loading, error } = useSelector((state) => state.assignments);
  const [selectedSubmission, setSelectedSubmission] = useState(null);
  const [feedback, setFeedback] = useState('');
  const [grade, setGrade] = useState('');

  useEffect(() => {
    dispatch(fetchSubmissionsForAssignment(id));
  }, [dispatch, id]);

  const handleGradeSubmission = (submissionId) => {
    dispatch(gradeSubmission({ submissionId, feedback, grade }))
      .then(() => {
        alert('Feedback and grade submitted!');
        setSelectedSubmission(null);
        setFeedback('');
        setGrade('');
      })
      .catch(() => alert('Submission failed. Please try again.'));
  };

  return (
    <div>
      <h1>Submissions for Assignment</h1>
      {loading && <p>Loading submissions...</p>}
      {error && <p>Error loading submissions: {error}</p>}
      <ul>
        {submissions.map((submission) => (
          <li key={submission.id}>
            <h3>Student: {submission.studentName}</h3>
            <p>Submission Date: {submission.submissionDate}</p>
            <button onClick={() => setSelectedSubmission(submission)}>Grade Submission</button>
          </li>
        ))}
      </ul>

      {selectedSubmission && (
        <div className="grading-form">
          <h2>Grade Submission for {selectedSubmission.studentName}</h2>
          <textarea
            value={feedback}
            onChange={(e) => setFeedback(e.target.value)}
            placeholder="Feedback"
          />
          <input
            type="text"
            value={grade}
            onChange={(e) => setGrade(e.target.value)}
            placeholder="Grade"
          />
          <button onClick={() => handleGradeSubmission(selectedSubmission.id)}>Submit Grade</button>
        </div>
      )}
    </div>
  );
}

export default AssignmentSubmissions;
