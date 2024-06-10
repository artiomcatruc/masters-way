import {SchemasProblemPopulatedResponse} from "src/apiAutogenerated";
import {Problem} from "src/model/businessModel/Problem";

/**
 * Problem converter params
 */
interface ProblemDoneConverterParams {

  /**
   * Problem DTO
   */
  problemDTO: SchemasProblemPopulatedResponse;

  /**
   * Job's way name
   */
  wayName: string;

  /**
   *Job's way uuid
   */
  wayUuid: string;
}

/**
 * Convert {@link ProblemDTO} to {@link Problem}
 */
export const ProblemDTOToProblem = (params: ProblemDoneConverterParams): Problem => {
  return new Problem({
    ...params.problemDTO,
    createdAt: new Date(params.problemDTO.createdAt),
    updatedAt: new Date(params.problemDTO.updatedAt),
    wayName: params.wayName,
    wayUuid: params.wayUuid,
  });
};
