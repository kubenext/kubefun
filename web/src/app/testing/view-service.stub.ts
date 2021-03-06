// Copyright (c) 2019 the Kubefun contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
//

export function viewServiceStubFactory() {
  return {
    titleAsText: () => '',
    viewTitleAsText: () => 'Just a title',
  };
}

export const viewServiceStub = viewServiceStubFactory();
