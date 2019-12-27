describe('Context', () => {
  before(() => {
    cy.visit('/');
  });

  it('has kubeconfig context', () => {
    cy.contains(' kubefun-temporary ').click();

    cy.get('.active').contains(' kubefun-temporary ');
  });
});
