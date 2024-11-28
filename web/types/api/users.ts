import z from "zod";

export const SingleUserResponseSchema = z.object({
    userUuid: z.string(),
    email: z.string(),
    firstName: z.string(),
    lastName: z.string(),
    hasPasswordSet: z.boolean(),
    role: z.string(),
    deactivatedTs: z.nullable(z.string()),
    createdAt: z.string(),
});
export type SingleUserResponse = z.infer<typeof SingleUserResponseSchema>;

export const ListAllUsersResponseSchema = z.array(SingleUserResponseSchema);
export type ListAllUsersResponse = z.infer<typeof ListAllUsersResponseSchema>;

export type UserRole = "Admin" | "User";

export interface CreateUserRequest {
    email: string;
    firstName: string;
    lastName: string;
    role: UserRole;
}

export interface UpdateUserRequest {
    firstName: string;
    lastName: string;
    role: UserRole;
}
