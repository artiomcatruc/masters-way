// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createCommentStmt, err = db.PrepareContext(ctx, createComment); err != nil {
		return nil, fmt.Errorf("error preparing query CreateComment: %w", err)
	}
	if q.createDayReportStmt, err = db.PrepareContext(ctx, createDayReport); err != nil {
		return nil, fmt.Errorf("error preparing query CreateDayReport: %w", err)
	}
	if q.createFavoriteUserStmt, err = db.PrepareContext(ctx, createFavoriteUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateFavoriteUser: %w", err)
	}
	if q.createFavoriteUserWayStmt, err = db.PrepareContext(ctx, createFavoriteUserWay); err != nil {
		return nil, fmt.Errorf("error preparing query CreateFavoriteUserWay: %w", err)
	}
	if q.createFormerMentorsWayStmt, err = db.PrepareContext(ctx, createFormerMentorsWay); err != nil {
		return nil, fmt.Errorf("error preparing query CreateFormerMentorsWay: %w", err)
	}
	if q.createFromUserMentoringRequestStmt, err = db.PrepareContext(ctx, createFromUserMentoringRequest); err != nil {
		return nil, fmt.Errorf("error preparing query CreateFromUserMentoringRequest: %w", err)
	}
	if q.createJobDoneStmt, err = db.PrepareContext(ctx, createJobDone); err != nil {
		return nil, fmt.Errorf("error preparing query CreateJobDone: %w", err)
	}
	if q.createJobDonesJobTagStmt, err = db.PrepareContext(ctx, createJobDonesJobTag); err != nil {
		return nil, fmt.Errorf("error preparing query CreateJobDonesJobTag: %w", err)
	}
	if q.createJobTagStmt, err = db.PrepareContext(ctx, createJobTag); err != nil {
		return nil, fmt.Errorf("error preparing query CreateJobTag: %w", err)
	}
	if q.createMetricStmt, err = db.PrepareContext(ctx, createMetric); err != nil {
		return nil, fmt.Errorf("error preparing query CreateMetric: %w", err)
	}
	if q.createPlanStmt, err = db.PrepareContext(ctx, createPlan); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePlan: %w", err)
	}
	if q.createPlansJobTagStmt, err = db.PrepareContext(ctx, createPlansJobTag); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePlansJobTag: %w", err)
	}
	if q.createProblemStmt, err = db.PrepareContext(ctx, createProblem); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProblem: %w", err)
	}
	if q.createProblemsJobTagStmt, err = db.PrepareContext(ctx, createProblemsJobTag); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProblemsJobTag: %w", err)
	}
	if q.createToUserMentoringRequestStmt, err = db.PrepareContext(ctx, createToUserMentoringRequest); err != nil {
		return nil, fmt.Errorf("error preparing query CreateToUserMentoringRequest: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.createUserTagStmt, err = db.PrepareContext(ctx, createUserTag); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUserTag: %w", err)
	}
	if q.createWayStmt, err = db.PrepareContext(ctx, createWay); err != nil {
		return nil, fmt.Errorf("error preparing query CreateWay: %w", err)
	}
	if q.createWayCollectionStmt, err = db.PrepareContext(ctx, createWayCollection); err != nil {
		return nil, fmt.Errorf("error preparing query CreateWayCollection: %w", err)
	}
	if q.createWayCollectionsWaysStmt, err = db.PrepareContext(ctx, createWayCollectionsWays); err != nil {
		return nil, fmt.Errorf("error preparing query CreateWayCollectionsWays: %w", err)
	}
	if q.createWayTagStmt, err = db.PrepareContext(ctx, createWayTag); err != nil {
		return nil, fmt.Errorf("error preparing query CreateWayTag: %w", err)
	}
	if q.deleteCommentStmt, err = db.PrepareContext(ctx, deleteComment); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteComment: %w", err)
	}
	if q.deleteFavoriteUserByIdsStmt, err = db.PrepareContext(ctx, deleteFavoriteUserByIds); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteFavoriteUserByIds: %w", err)
	}
	if q.deleteFavoriteUserWayByIdsStmt, err = db.PrepareContext(ctx, deleteFavoriteUserWayByIds); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteFavoriteUserWayByIds: %w", err)
	}
	if q.deleteFromUserMentoringRequestByIdsStmt, err = db.PrepareContext(ctx, deleteFromUserMentoringRequestByIds); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteFromUserMentoringRequestByIds: %w", err)
	}
	if q.deleteJobDoneStmt, err = db.PrepareContext(ctx, deleteJobDone); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteJobDone: %w", err)
	}
	if q.deleteJobDonesJobTagByJobDoneIdStmt, err = db.PrepareContext(ctx, deleteJobDonesJobTagByJobDoneId); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteJobDonesJobTagByJobDoneId: %w", err)
	}
	if q.deleteJobTagByStmt, err = db.PrepareContext(ctx, deleteJobTagBy); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteJobTagBy: %w", err)
	}
	if q.deleteMetricStmt, err = db.PrepareContext(ctx, deleteMetric); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteMetric: %w", err)
	}
	if q.deletePlanStmt, err = db.PrepareContext(ctx, deletePlan); err != nil {
		return nil, fmt.Errorf("error preparing query DeletePlan: %w", err)
	}
	if q.deletePlansJobTagByIdsStmt, err = db.PrepareContext(ctx, deletePlansJobTagByIds); err != nil {
		return nil, fmt.Errorf("error preparing query DeletePlansJobTagByIds: %w", err)
	}
	if q.deleteProblemStmt, err = db.PrepareContext(ctx, deleteProblem); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProblem: %w", err)
	}
	if q.deleteProblemsJobTagByIdsStmt, err = db.PrepareContext(ctx, deleteProblemsJobTagByIds); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProblemsJobTagByIds: %w", err)
	}
	if q.deleteToUserMentoringRequestByIdsStmt, err = db.PrepareContext(ctx, deleteToUserMentoringRequestByIds); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteToUserMentoringRequestByIds: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, deleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.deleteUserTagStmt, err = db.PrepareContext(ctx, deleteUserTag); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUserTag: %w", err)
	}
	if q.deleteWayStmt, err = db.PrepareContext(ctx, deleteWay); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteWay: %w", err)
	}
	if q.deleteWayCollectionStmt, err = db.PrepareContext(ctx, deleteWayCollection); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteWayCollection: %w", err)
	}
	if q.deleteWayCollectionsWaysByIdsStmt, err = db.PrepareContext(ctx, deleteWayCollectionsWaysByIds); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteWayCollectionsWaysByIds: %w", err)
	}
	if q.deleteWayTagStmt, err = db.PrepareContext(ctx, deleteWayTag); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteWayTag: %w", err)
	}
	if q.getJobDonesJoinJobTagsStmt, err = db.PrepareContext(ctx, getJobDonesJoinJobTags); err != nil {
		return nil, fmt.Errorf("error preparing query GetJobDonesJoinJobTags: %w", err)
	}
	if q.getListCommentsByDayReportIdStmt, err = db.PrepareContext(ctx, getListCommentsByDayReportId); err != nil {
		return nil, fmt.Errorf("error preparing query GetListCommentsByDayReportId: %w", err)
	}
	if q.getListDayReportsByWayUuidStmt, err = db.PrepareContext(ctx, getListDayReportsByWayUuid); err != nil {
		return nil, fmt.Errorf("error preparing query GetListDayReportsByWayUuid: %w", err)
	}
	if q.getListJobTagsByWayUuidStmt, err = db.PrepareContext(ctx, getListJobTagsByWayUuid); err != nil {
		return nil, fmt.Errorf("error preparing query GetListJobTagsByWayUuid: %w", err)
	}
	if q.getListJobsDoneByDayReportIdStmt, err = db.PrepareContext(ctx, getListJobsDoneByDayReportId); err != nil {
		return nil, fmt.Errorf("error preparing query GetListJobsDoneByDayReportId: %w", err)
	}
	if q.getListMetricsByWayUuidStmt, err = db.PrepareContext(ctx, getListMetricsByWayUuid); err != nil {
		return nil, fmt.Errorf("error preparing query GetListMetricsByWayUuid: %w", err)
	}
	if q.getListPlansByDayReportIdStmt, err = db.PrepareContext(ctx, getListPlansByDayReportId); err != nil {
		return nil, fmt.Errorf("error preparing query GetListPlansByDayReportId: %w", err)
	}
	if q.getListProblemsByDayReportIdStmt, err = db.PrepareContext(ctx, getListProblemsByDayReportId); err != nil {
		return nil, fmt.Errorf("error preparing query GetListProblemsByDayReportId: %w", err)
	}
	if q.getListUserTagsByUserIdStmt, err = db.PrepareContext(ctx, getListUserTagsByUserId); err != nil {
		return nil, fmt.Errorf("error preparing query GetListUserTagsByUserId: %w", err)
	}
	if q.getListWayCollectionsByUserIdStmt, err = db.PrepareContext(ctx, getListWayCollectionsByUserId); err != nil {
		return nil, fmt.Errorf("error preparing query GetListWayCollectionsByUserId: %w", err)
	}
	if q.getListWayTagsByWayIdStmt, err = db.PrepareContext(ctx, getListWayTagsByWayId); err != nil {
		return nil, fmt.Errorf("error preparing query GetListWayTagsByWayId: %w", err)
	}
	if q.getPlansJoinJobTagsStmt, err = db.PrepareContext(ctx, getPlansJoinJobTags); err != nil {
		return nil, fmt.Errorf("error preparing query GetPlansJoinJobTags: %w", err)
	}
	if q.getProblemsJoinJobTagsStmt, err = db.PrepareContext(ctx, getProblemsJoinJobTags); err != nil {
		return nil, fmt.Errorf("error preparing query GetProblemsJoinJobTags: %w", err)
	}
	if q.getUserByIdStmt, err = db.PrepareContext(ctx, getUserById); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserById: %w", err)
	}
	if q.getWayByIdStmt, err = db.PrepareContext(ctx, getWayById); err != nil {
		return nil, fmt.Errorf("error preparing query GetWayById: %w", err)
	}
	if q.listUsersStmt, err = db.PrepareContext(ctx, listUsers); err != nil {
		return nil, fmt.Errorf("error preparing query ListUsers: %w", err)
	}
	if q.listWaysStmt, err = db.PrepareContext(ctx, listWays); err != nil {
		return nil, fmt.Errorf("error preparing query ListWays: %w", err)
	}
	if q.updateCommentStmt, err = db.PrepareContext(ctx, updateComment); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateComment: %w", err)
	}
	if q.updateDayReportStmt, err = db.PrepareContext(ctx, updateDayReport); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateDayReport: %w", err)
	}
	if q.updateJobDoneStmt, err = db.PrepareContext(ctx, updateJobDone); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateJobDone: %w", err)
	}
	if q.updateJobTagStmt, err = db.PrepareContext(ctx, updateJobTag); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateJobTag: %w", err)
	}
	if q.updateMetricStmt, err = db.PrepareContext(ctx, updateMetric); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateMetric: %w", err)
	}
	if q.updatePlanStmt, err = db.PrepareContext(ctx, updatePlan); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePlan: %w", err)
	}
	if q.updateProblemStmt, err = db.PrepareContext(ctx, updateProblem); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateProblem: %w", err)
	}
	if q.updateUserStmt, err = db.PrepareContext(ctx, updateUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUser: %w", err)
	}
	if q.updateUserTagStmt, err = db.PrepareContext(ctx, updateUserTag); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUserTag: %w", err)
	}
	if q.updateWayStmt, err = db.PrepareContext(ctx, updateWay); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateWay: %w", err)
	}
	if q.updateWayCollectionStmt, err = db.PrepareContext(ctx, updateWayCollection); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateWayCollection: %w", err)
	}
	if q.updateWayTagStmt, err = db.PrepareContext(ctx, updateWayTag); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateWayTag: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createCommentStmt != nil {
		if cerr := q.createCommentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createCommentStmt: %w", cerr)
		}
	}
	if q.createDayReportStmt != nil {
		if cerr := q.createDayReportStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createDayReportStmt: %w", cerr)
		}
	}
	if q.createFavoriteUserStmt != nil {
		if cerr := q.createFavoriteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createFavoriteUserStmt: %w", cerr)
		}
	}
	if q.createFavoriteUserWayStmt != nil {
		if cerr := q.createFavoriteUserWayStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createFavoriteUserWayStmt: %w", cerr)
		}
	}
	if q.createFormerMentorsWayStmt != nil {
		if cerr := q.createFormerMentorsWayStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createFormerMentorsWayStmt: %w", cerr)
		}
	}
	if q.createFromUserMentoringRequestStmt != nil {
		if cerr := q.createFromUserMentoringRequestStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createFromUserMentoringRequestStmt: %w", cerr)
		}
	}
	if q.createJobDoneStmt != nil {
		if cerr := q.createJobDoneStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createJobDoneStmt: %w", cerr)
		}
	}
	if q.createJobDonesJobTagStmt != nil {
		if cerr := q.createJobDonesJobTagStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createJobDonesJobTagStmt: %w", cerr)
		}
	}
	if q.createJobTagStmt != nil {
		if cerr := q.createJobTagStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createJobTagStmt: %w", cerr)
		}
	}
	if q.createMetricStmt != nil {
		if cerr := q.createMetricStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createMetricStmt: %w", cerr)
		}
	}
	if q.createPlanStmt != nil {
		if cerr := q.createPlanStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPlanStmt: %w", cerr)
		}
	}
	if q.createPlansJobTagStmt != nil {
		if cerr := q.createPlansJobTagStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPlansJobTagStmt: %w", cerr)
		}
	}
	if q.createProblemStmt != nil {
		if cerr := q.createProblemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProblemStmt: %w", cerr)
		}
	}
	if q.createProblemsJobTagStmt != nil {
		if cerr := q.createProblemsJobTagStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProblemsJobTagStmt: %w", cerr)
		}
	}
	if q.createToUserMentoringRequestStmt != nil {
		if cerr := q.createToUserMentoringRequestStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createToUserMentoringRequestStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.createUserTagStmt != nil {
		if cerr := q.createUserTagStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserTagStmt: %w", cerr)
		}
	}
	if q.createWayStmt != nil {
		if cerr := q.createWayStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createWayStmt: %w", cerr)
		}
	}
	if q.createWayCollectionStmt != nil {
		if cerr := q.createWayCollectionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createWayCollectionStmt: %w", cerr)
		}
	}
	if q.createWayCollectionsWaysStmt != nil {
		if cerr := q.createWayCollectionsWaysStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createWayCollectionsWaysStmt: %w", cerr)
		}
	}
	if q.createWayTagStmt != nil {
		if cerr := q.createWayTagStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createWayTagStmt: %w", cerr)
		}
	}
	if q.deleteCommentStmt != nil {
		if cerr := q.deleteCommentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteCommentStmt: %w", cerr)
		}
	}
	if q.deleteFavoriteUserByIdsStmt != nil {
		if cerr := q.deleteFavoriteUserByIdsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteFavoriteUserByIdsStmt: %w", cerr)
		}
	}
	if q.deleteFavoriteUserWayByIdsStmt != nil {
		if cerr := q.deleteFavoriteUserWayByIdsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteFavoriteUserWayByIdsStmt: %w", cerr)
		}
	}
	if q.deleteFromUserMentoringRequestByIdsStmt != nil {
		if cerr := q.deleteFromUserMentoringRequestByIdsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteFromUserMentoringRequestByIdsStmt: %w", cerr)
		}
	}
	if q.deleteJobDoneStmt != nil {
		if cerr := q.deleteJobDoneStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteJobDoneStmt: %w", cerr)
		}
	}
	if q.deleteJobDonesJobTagByJobDoneIdStmt != nil {
		if cerr := q.deleteJobDonesJobTagByJobDoneIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteJobDonesJobTagByJobDoneIdStmt: %w", cerr)
		}
	}
	if q.deleteJobTagByStmt != nil {
		if cerr := q.deleteJobTagByStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteJobTagByStmt: %w", cerr)
		}
	}
	if q.deleteMetricStmt != nil {
		if cerr := q.deleteMetricStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteMetricStmt: %w", cerr)
		}
	}
	if q.deletePlanStmt != nil {
		if cerr := q.deletePlanStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deletePlanStmt: %w", cerr)
		}
	}
	if q.deletePlansJobTagByIdsStmt != nil {
		if cerr := q.deletePlansJobTagByIdsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deletePlansJobTagByIdsStmt: %w", cerr)
		}
	}
	if q.deleteProblemStmt != nil {
		if cerr := q.deleteProblemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteProblemStmt: %w", cerr)
		}
	}
	if q.deleteProblemsJobTagByIdsStmt != nil {
		if cerr := q.deleteProblemsJobTagByIdsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteProblemsJobTagByIdsStmt: %w", cerr)
		}
	}
	if q.deleteToUserMentoringRequestByIdsStmt != nil {
		if cerr := q.deleteToUserMentoringRequestByIdsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteToUserMentoringRequestByIdsStmt: %w", cerr)
		}
	}
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
		}
	}
	if q.deleteUserTagStmt != nil {
		if cerr := q.deleteUserTagStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserTagStmt: %w", cerr)
		}
	}
	if q.deleteWayStmt != nil {
		if cerr := q.deleteWayStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteWayStmt: %w", cerr)
		}
	}
	if q.deleteWayCollectionStmt != nil {
		if cerr := q.deleteWayCollectionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteWayCollectionStmt: %w", cerr)
		}
	}
	if q.deleteWayCollectionsWaysByIdsStmt != nil {
		if cerr := q.deleteWayCollectionsWaysByIdsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteWayCollectionsWaysByIdsStmt: %w", cerr)
		}
	}
	if q.deleteWayTagStmt != nil {
		if cerr := q.deleteWayTagStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteWayTagStmt: %w", cerr)
		}
	}
	if q.getJobDonesJoinJobTagsStmt != nil {
		if cerr := q.getJobDonesJoinJobTagsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getJobDonesJoinJobTagsStmt: %w", cerr)
		}
	}
	if q.getListCommentsByDayReportIdStmt != nil {
		if cerr := q.getListCommentsByDayReportIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getListCommentsByDayReportIdStmt: %w", cerr)
		}
	}
	if q.getListDayReportsByWayUuidStmt != nil {
		if cerr := q.getListDayReportsByWayUuidStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getListDayReportsByWayUuidStmt: %w", cerr)
		}
	}
	if q.getListJobTagsByWayUuidStmt != nil {
		if cerr := q.getListJobTagsByWayUuidStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getListJobTagsByWayUuidStmt: %w", cerr)
		}
	}
	if q.getListJobsDoneByDayReportIdStmt != nil {
		if cerr := q.getListJobsDoneByDayReportIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getListJobsDoneByDayReportIdStmt: %w", cerr)
		}
	}
	if q.getListMetricsByWayUuidStmt != nil {
		if cerr := q.getListMetricsByWayUuidStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getListMetricsByWayUuidStmt: %w", cerr)
		}
	}
	if q.getListPlansByDayReportIdStmt != nil {
		if cerr := q.getListPlansByDayReportIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getListPlansByDayReportIdStmt: %w", cerr)
		}
	}
	if q.getListProblemsByDayReportIdStmt != nil {
		if cerr := q.getListProblemsByDayReportIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getListProblemsByDayReportIdStmt: %w", cerr)
		}
	}
	if q.getListUserTagsByUserIdStmt != nil {
		if cerr := q.getListUserTagsByUserIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getListUserTagsByUserIdStmt: %w", cerr)
		}
	}
	if q.getListWayCollectionsByUserIdStmt != nil {
		if cerr := q.getListWayCollectionsByUserIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getListWayCollectionsByUserIdStmt: %w", cerr)
		}
	}
	if q.getListWayTagsByWayIdStmt != nil {
		if cerr := q.getListWayTagsByWayIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getListWayTagsByWayIdStmt: %w", cerr)
		}
	}
	if q.getPlansJoinJobTagsStmt != nil {
		if cerr := q.getPlansJoinJobTagsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPlansJoinJobTagsStmt: %w", cerr)
		}
	}
	if q.getProblemsJoinJobTagsStmt != nil {
		if cerr := q.getProblemsJoinJobTagsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProblemsJoinJobTagsStmt: %w", cerr)
		}
	}
	if q.getUserByIdStmt != nil {
		if cerr := q.getUserByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByIdStmt: %w", cerr)
		}
	}
	if q.getWayByIdStmt != nil {
		if cerr := q.getWayByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getWayByIdStmt: %w", cerr)
		}
	}
	if q.listUsersStmt != nil {
		if cerr := q.listUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listUsersStmt: %w", cerr)
		}
	}
	if q.listWaysStmt != nil {
		if cerr := q.listWaysStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listWaysStmt: %w", cerr)
		}
	}
	if q.updateCommentStmt != nil {
		if cerr := q.updateCommentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateCommentStmt: %w", cerr)
		}
	}
	if q.updateDayReportStmt != nil {
		if cerr := q.updateDayReportStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateDayReportStmt: %w", cerr)
		}
	}
	if q.updateJobDoneStmt != nil {
		if cerr := q.updateJobDoneStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateJobDoneStmt: %w", cerr)
		}
	}
	if q.updateJobTagStmt != nil {
		if cerr := q.updateJobTagStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateJobTagStmt: %w", cerr)
		}
	}
	if q.updateMetricStmt != nil {
		if cerr := q.updateMetricStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateMetricStmt: %w", cerr)
		}
	}
	if q.updatePlanStmt != nil {
		if cerr := q.updatePlanStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePlanStmt: %w", cerr)
		}
	}
	if q.updateProblemStmt != nil {
		if cerr := q.updateProblemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateProblemStmt: %w", cerr)
		}
	}
	if q.updateUserStmt != nil {
		if cerr := q.updateUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserStmt: %w", cerr)
		}
	}
	if q.updateUserTagStmt != nil {
		if cerr := q.updateUserTagStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserTagStmt: %w", cerr)
		}
	}
	if q.updateWayStmt != nil {
		if cerr := q.updateWayStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateWayStmt: %w", cerr)
		}
	}
	if q.updateWayCollectionStmt != nil {
		if cerr := q.updateWayCollectionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateWayCollectionStmt: %w", cerr)
		}
	}
	if q.updateWayTagStmt != nil {
		if cerr := q.updateWayTagStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateWayTagStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                      DBTX
	tx                                      *sql.Tx
	createCommentStmt                       *sql.Stmt
	createDayReportStmt                     *sql.Stmt
	createFavoriteUserStmt                  *sql.Stmt
	createFavoriteUserWayStmt               *sql.Stmt
	createFormerMentorsWayStmt              *sql.Stmt
	createFromUserMentoringRequestStmt      *sql.Stmt
	createJobDoneStmt                       *sql.Stmt
	createJobDonesJobTagStmt                *sql.Stmt
	createJobTagStmt                        *sql.Stmt
	createMetricStmt                        *sql.Stmt
	createPlanStmt                          *sql.Stmt
	createPlansJobTagStmt                   *sql.Stmt
	createProblemStmt                       *sql.Stmt
	createProblemsJobTagStmt                *sql.Stmt
	createToUserMentoringRequestStmt        *sql.Stmt
	createUserStmt                          *sql.Stmt
	createUserTagStmt                       *sql.Stmt
	createWayStmt                           *sql.Stmt
	createWayCollectionStmt                 *sql.Stmt
	createWayCollectionsWaysStmt            *sql.Stmt
	createWayTagStmt                        *sql.Stmt
	deleteCommentStmt                       *sql.Stmt
	deleteFavoriteUserByIdsStmt             *sql.Stmt
	deleteFavoriteUserWayByIdsStmt          *sql.Stmt
	deleteFromUserMentoringRequestByIdsStmt *sql.Stmt
	deleteJobDoneStmt                       *sql.Stmt
	deleteJobDonesJobTagByJobDoneIdStmt     *sql.Stmt
	deleteJobTagByStmt                      *sql.Stmt
	deleteMetricStmt                        *sql.Stmt
	deletePlanStmt                          *sql.Stmt
	deletePlansJobTagByIdsStmt              *sql.Stmt
	deleteProblemStmt                       *sql.Stmt
	deleteProblemsJobTagByIdsStmt           *sql.Stmt
	deleteToUserMentoringRequestByIdsStmt   *sql.Stmt
	deleteUserStmt                          *sql.Stmt
	deleteUserTagStmt                       *sql.Stmt
	deleteWayStmt                           *sql.Stmt
	deleteWayCollectionStmt                 *sql.Stmt
	deleteWayCollectionsWaysByIdsStmt       *sql.Stmt
	deleteWayTagStmt                        *sql.Stmt
	getJobDonesJoinJobTagsStmt              *sql.Stmt
	getListCommentsByDayReportIdStmt        *sql.Stmt
	getListDayReportsByWayUuidStmt          *sql.Stmt
	getListJobTagsByWayUuidStmt             *sql.Stmt
	getListJobsDoneByDayReportIdStmt        *sql.Stmt
	getListMetricsByWayUuidStmt             *sql.Stmt
	getListPlansByDayReportIdStmt           *sql.Stmt
	getListProblemsByDayReportIdStmt        *sql.Stmt
	getListUserTagsByUserIdStmt             *sql.Stmt
	getListWayCollectionsByUserIdStmt       *sql.Stmt
	getListWayTagsByWayIdStmt               *sql.Stmt
	getPlansJoinJobTagsStmt                 *sql.Stmt
	getProblemsJoinJobTagsStmt              *sql.Stmt
	getUserByIdStmt                         *sql.Stmt
	getWayByIdStmt                          *sql.Stmt
	listUsersStmt                           *sql.Stmt
	listWaysStmt                            *sql.Stmt
	updateCommentStmt                       *sql.Stmt
	updateDayReportStmt                     *sql.Stmt
	updateJobDoneStmt                       *sql.Stmt
	updateJobTagStmt                        *sql.Stmt
	updateMetricStmt                        *sql.Stmt
	updatePlanStmt                          *sql.Stmt
	updateProblemStmt                       *sql.Stmt
	updateUserStmt                          *sql.Stmt
	updateUserTagStmt                       *sql.Stmt
	updateWayStmt                           *sql.Stmt
	updateWayCollectionStmt                 *sql.Stmt
	updateWayTagStmt                        *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                      tx,
		tx:                                      tx,
		createCommentStmt:                       q.createCommentStmt,
		createDayReportStmt:                     q.createDayReportStmt,
		createFavoriteUserStmt:                  q.createFavoriteUserStmt,
		createFavoriteUserWayStmt:               q.createFavoriteUserWayStmt,
		createFormerMentorsWayStmt:              q.createFormerMentorsWayStmt,
		createFromUserMentoringRequestStmt:      q.createFromUserMentoringRequestStmt,
		createJobDoneStmt:                       q.createJobDoneStmt,
		createJobDonesJobTagStmt:                q.createJobDonesJobTagStmt,
		createJobTagStmt:                        q.createJobTagStmt,
		createMetricStmt:                        q.createMetricStmt,
		createPlanStmt:                          q.createPlanStmt,
		createPlansJobTagStmt:                   q.createPlansJobTagStmt,
		createProblemStmt:                       q.createProblemStmt,
		createProblemsJobTagStmt:                q.createProblemsJobTagStmt,
		createToUserMentoringRequestStmt:        q.createToUserMentoringRequestStmt,
		createUserStmt:                          q.createUserStmt,
		createUserTagStmt:                       q.createUserTagStmt,
		createWayStmt:                           q.createWayStmt,
		createWayCollectionStmt:                 q.createWayCollectionStmt,
		createWayCollectionsWaysStmt:            q.createWayCollectionsWaysStmt,
		createWayTagStmt:                        q.createWayTagStmt,
		deleteCommentStmt:                       q.deleteCommentStmt,
		deleteFavoriteUserByIdsStmt:             q.deleteFavoriteUserByIdsStmt,
		deleteFavoriteUserWayByIdsStmt:          q.deleteFavoriteUserWayByIdsStmt,
		deleteFromUserMentoringRequestByIdsStmt: q.deleteFromUserMentoringRequestByIdsStmt,
		deleteJobDoneStmt:                       q.deleteJobDoneStmt,
		deleteJobDonesJobTagByJobDoneIdStmt:     q.deleteJobDonesJobTagByJobDoneIdStmt,
		deleteJobTagByStmt:                      q.deleteJobTagByStmt,
		deleteMetricStmt:                        q.deleteMetricStmt,
		deletePlanStmt:                          q.deletePlanStmt,
		deletePlansJobTagByIdsStmt:              q.deletePlansJobTagByIdsStmt,
		deleteProblemStmt:                       q.deleteProblemStmt,
		deleteProblemsJobTagByIdsStmt:           q.deleteProblemsJobTagByIdsStmt,
		deleteToUserMentoringRequestByIdsStmt:   q.deleteToUserMentoringRequestByIdsStmt,
		deleteUserStmt:                          q.deleteUserStmt,
		deleteUserTagStmt:                       q.deleteUserTagStmt,
		deleteWayStmt:                           q.deleteWayStmt,
		deleteWayCollectionStmt:                 q.deleteWayCollectionStmt,
		deleteWayCollectionsWaysByIdsStmt:       q.deleteWayCollectionsWaysByIdsStmt,
		deleteWayTagStmt:                        q.deleteWayTagStmt,
		getJobDonesJoinJobTagsStmt:              q.getJobDonesJoinJobTagsStmt,
		getListCommentsByDayReportIdStmt:        q.getListCommentsByDayReportIdStmt,
		getListDayReportsByWayUuidStmt:          q.getListDayReportsByWayUuidStmt,
		getListJobTagsByWayUuidStmt:             q.getListJobTagsByWayUuidStmt,
		getListJobsDoneByDayReportIdStmt:        q.getListJobsDoneByDayReportIdStmt,
		getListMetricsByWayUuidStmt:             q.getListMetricsByWayUuidStmt,
		getListPlansByDayReportIdStmt:           q.getListPlansByDayReportIdStmt,
		getListProblemsByDayReportIdStmt:        q.getListProblemsByDayReportIdStmt,
		getListUserTagsByUserIdStmt:             q.getListUserTagsByUserIdStmt,
		getListWayCollectionsByUserIdStmt:       q.getListWayCollectionsByUserIdStmt,
		getListWayTagsByWayIdStmt:               q.getListWayTagsByWayIdStmt,
		getPlansJoinJobTagsStmt:                 q.getPlansJoinJobTagsStmt,
		getProblemsJoinJobTagsStmt:              q.getProblemsJoinJobTagsStmt,
		getUserByIdStmt:                         q.getUserByIdStmt,
		getWayByIdStmt:                          q.getWayByIdStmt,
		listUsersStmt:                           q.listUsersStmt,
		listWaysStmt:                            q.listWaysStmt,
		updateCommentStmt:                       q.updateCommentStmt,
		updateDayReportStmt:                     q.updateDayReportStmt,
		updateJobDoneStmt:                       q.updateJobDoneStmt,
		updateJobTagStmt:                        q.updateJobTagStmt,
		updateMetricStmt:                        q.updateMetricStmt,
		updatePlanStmt:                          q.updatePlanStmt,
		updateProblemStmt:                       q.updateProblemStmt,
		updateUserStmt:                          q.updateUserStmt,
		updateUserTagStmt:                       q.updateUserTagStmt,
		updateWayStmt:                           q.updateWayStmt,
		updateWayCollectionStmt:                 q.updateWayCollectionStmt,
		updateWayTagStmt:                        q.updateWayTagStmt,
	}
}