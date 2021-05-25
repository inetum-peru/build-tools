export enum ValidatePrompNames {
  'name' = 'nameValidate'
}

export enum BuilderPrompNames {
  'name' = 'nameBuilder'
}

export type AnswersValidate = { [P in ValidatePrompNames]: string }
export type AnswersBuilder = { [P in BuilderPrompNames]: string }
