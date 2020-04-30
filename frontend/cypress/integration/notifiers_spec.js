/// <reference types="cypress" />

import "../support/commands"

context('Notifier Tests', () => {


  beforeEach(() => {
    cy.restoreLocalStorageCache();
  });

  afterEach(() => {
    cy.saveLocalStorageCache();
  });

  it('should Login', () => {
    cy.visit('/login')
    cy.get('#username').clear().type('admin')
    cy.get('#password').clear().type('admin')
    cy.get('button[type="submit"]').click()

    cy.get('.navbar-brand').should('contain', 'Statping')
    cy.getCookies()

    cy.getCookies().should('have.length', 1)
  })

  // uzrwstmtd69hi4wgzsj27q2v29mtpu
  it('should confirm notifiers are installed', () => {
    cy.visit('/dashboard/settings')
    cy.get('#notifiers_tabs > a').should('have.length', 10)

    cy.get('#api_secret').should('not.have.value', '')
  })

  // it('should test and save notifier', () => {
  //   cy.visit('/dashboard/settings')
  //   cy.get('#notifiers_tabs > a').should('have.length', 10)
  //   cy.get('#notifiers_tabs > #v-pills-command-tab').click()
  //
  //   cy.get('#v-pills-command-tab > .form-control').eq(0).clear().type('/bin/echo')
  //   cy.get('#v-pills-command-tab > .form-control').eq(1).clear().type('"success"')
  //   cy.get('#v-pills-command-tab > .form-control').eq(2).clear().type('"failure"')
  //
  //   cy.get('#v-pills-command-tab').find(".save-notifier").click()
  //   cy.get('#v-pills-command-tab').find(".test-notifier").click()
  // })

})
