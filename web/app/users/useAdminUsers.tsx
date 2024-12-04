"use client";

import { Admin } from "@/api/admin";
import { useGetAllUsers } from "@/api/hooks";
import { useNotificationsContext } from "@/context/NotificationsContext/NotificationsContext";
import { ApplicationError } from "@/types/api/error";
import type { CreateUserRequest, SingleUserResponse, UpdateUserRequest } from "@/types/api/users";
import type React from "react";
import { useState } from "react";
import { createContext, useContext } from "react";

const AdminUsersContext = createContext<UseAdminUsers | undefined>(undefined);

const useAdminUsersHook = (): UseAdminUsers => {
    const { addNotification } = useNotificationsContext();
    const [isUpdating, setIsUpdating] = useState(false);
    const [updateError, setUpdateError] = useState<ApplicationError | undefined>(undefined);
    const {
        mutate: reloadUsers,
        data: allUsers,
        error: allUsersError,
        isLoading: allUsersIsLoading,
    } = useGetAllUsers();

    const addUser = async (newUser: CreateUserRequest, onSuccess: () => void) => {
        setUpdateError(undefined);
        setIsUpdating(true);
        const [res, err] = await Admin.CreateUser(newUser);
        if (err) {
            setUpdateError(new ApplicationError(err));
        } else {
            onSuccess();
            addNotification({
                severity: "success",
                title: "User added",
                message: `The user for ${newUser.firstName} ${newUser.lastName} with role ${newUser.role} has been added successfully`,
            });
        }
        await reloadUsers();
        setIsUpdating(false);
    };

    const updateUser = async (userUuid: string, updatedUser: UpdateUserRequest) => {
        setUpdateError(undefined);
        setIsUpdating(true);
        const [res, err] = await Admin.UpdateUser(userUuid, updatedUser);
        if (err) {
            setUpdateError(new ApplicationError(err));
        } else {
            addNotification({
                severity: "success",
                title: "User updated",
                message: `The user ${userUuid} has been updated successfully`,
            });
        }
        await reloadUsers();
        setIsUpdating(false);
    };

    const deactivateUser = async (userUuid: string) => {
        setUpdateError(undefined);
        setIsUpdating(true);
        const [res, err] = await Admin.DeactivateUser(userUuid);
        if (err) {
            setUpdateError(new ApplicationError(err));
        } else {
            addNotification({
                severity: "success",
                title: "User deactivated",
                message: `The user ${userUuid} has been deactivated`,
            });
        }
        await reloadUsers();
        setIsUpdating(false);
    };

    const reactivateUser = async (userUuid: string) => {
        setUpdateError(undefined);
        setIsUpdating(true);
        const [res, err] = await Admin.ReactivateUser(userUuid);
        if (err) {
            setUpdateError(new ApplicationError(err));
        } else {
            addNotification({
                severity: "success",
                title: "User reactivated",
                message: `The user ${userUuid} has been reactivated`,
            });
        }
        await reloadUsers();
        setIsUpdating(false);
    };

    return {
        allUsers,
        allUsersIsLoading,
        isUpdating,
        addUser,
        updateUser,
        deactivateUser,
        reactivateUser,
        errors: [allUsersError, updateError],
    };
};

export const AdminUsersContextProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
    const hook = useAdminUsersHook();

    return <AdminUsersContext.Provider value={hook}>{children}</AdminUsersContext.Provider>;
};

const useAdminUsers = () => {
    const context = useContext(AdminUsersContext);
    if (!context) {
        throw new Error("useAdminUsers must be used within an AdminUsersContextProvider");
    }
    return context;
};

export default useAdminUsers;

export interface UseAdminUsers {
    allUsers?: SingleUserResponse[];
    allUsersIsLoading: boolean;
    isUpdating: boolean;
    addUser: (newUser: CreateUserRequest, onSuccess: () => void) => void;
    updateUser: (userUuid: string, updatedUser: UpdateUserRequest) => void;
    deactivateUser: (userUuid: string) => void;
    reactivateUser: (userUuid: string) => void;
    errors: (ApplicationError | undefined)[];
}
