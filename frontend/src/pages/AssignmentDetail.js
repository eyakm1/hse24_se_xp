import React, { useEffect } from 'react';
import { useParams, Link } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { fetchAssignmentDetail, fetchSubmissionStatus } from '../store/slices/assignmentsSlice';

function AssignmentDetail() {
  const { id } = useParams();
  const dispatch = useDispatch();
  const { detail, loading, error, submissionStatus } = useSelector((state) => state.assignments);

  useEffect(() => {
    dispatch(fetchAssignmentDetail(id));
    dispatch(fetchSubmissionStatus(id));
  }, [dispatch, id]);

  return (
    <div>
      {loading && <p>Loading assignment details...</p>}
      {error && <p>Error loading details: {error}</p>}
      {detail && (
        <div>
          <h1>{detail.title}</h1>
          <p>{detail.description}</p>
          <p>Due Date: {detail.dueDate}</p>
          <Link to={`/assignments/${id}/submit`}>Submit Assignment</Link>

          {submissionStatus && (
            <div className="submission-status">
              <h2>Submission Status</h2>
              <p>Status: {submissionStatus.status}</p>
              {submissionStatus.submissionDate && <p>Submitted on: {submissionStatus.submissionDate}</p>}
              {submissionStatus.feedback && <p>Feedback: {submissionStatus.feedback}</p>}
              {submissionStatus.grade && <p>Grade: {submissionStatus.grade}</p>}
            </div>
          )}
        </div>
      )}
    </div>
  );
}

export default AssignmentDetail;
